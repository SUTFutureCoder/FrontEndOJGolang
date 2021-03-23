package models

import (
	"log"
	"strconv"
	"strings"
)

// ContestUserMap 比赛用户关联表
type ContestUserMap struct {
	Model
	// ContestId 比赛Id
	ContestId uint64 `json:"contest_id"`
}

const TABLE_CONTEST_USER_MAP = "contest_user_map"

func (c *ContestUserMap) Insert() (int64, error) {
	stmt, err := DB.Prepare("INSERT INTO contest_user_map (contest_id, status, creator_id, creator, create_time) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Printf("[ERROR] database exec error input[%v] err[%v]", c, err)
		return 0, err
	}
	defer stmt.Close()
	ret, err := stmt.Exec(
		c.ContestId,
		c.CreatorId,
		c.Status,
		c.Creator,
		c.CreateTime,
	)
	if err != nil || ret == nil {
		return 0, err
	}
	return ret.LastInsertId()
}

func (c *ContestUserMap) GetList(page Pager, status int) ([]*ContestUserMap, error) {
	var contestUserMaps []*ContestUserMap
	stmt, rows, err := GetByPager("SELECT id, contest_id, status, creator_id, creator, create_time, update_time FROM contest_user_map", page, status)
	defer stmt.Close()
	if err != nil {
		log.Printf("get contest user list from db error [%v]", err)
		return nil, err
	}

	if rows == nil {
		return nil, err
	}
	for rows.Next() {
		c := &ContestUserMap{}
		err = rows.Scan(
			&c.ID, &c.ContestId, &c.Status, &c.CreatorId, &c.Creator, &c.CreateTime, &c.UpdateTime,
		)
		contestUserMaps = append(contestUserMaps, c)
	}
	return contestUserMaps, err
}

func (c *ContestUserMap) GetByContestIds(contestIds []interface{}, status int) []*ContestUserMap {
	var contestUserMaps []*ContestUserMap
	if len(contestIds) == 0 {
		return contestUserMaps
	}

	query := "SELECT contest_id, creator_id, creator, status FROM contest_user_map WHERE contest_id IN (?"+strings.Repeat(",?", len(contestIds)-1)+")"
	if status != STATUS_ALL {
		contestIds = append(contestIds, status)
		query += " AND status=?"
	}
	rows, err := DB.Query(query, contestIds...)
	defer rows.Close()
	if err != nil {
		log.Printf("get contest user list by ids error [%v]\n", err)
		return contestUserMaps
	}
	for rows.Next() {
		c := &ContestUserMap{}
		err = rows.Scan(
			&c.ContestId, &c.CreatorId, &c.Creator, &c.Status,
		)
		if err != nil {
			log.Printf("scan lab list by ids ")
			return contestUserMaps
		}
		contestUserMaps = append(contestUserMaps, c)
	}
	return contestUserMaps
}

func (c *ContestUserMap) GetIdListByContestIds(contestIds []interface{}, status int) ([]*ContestUserMap, []interface{}) {
	contestUserMap := c.GetByContestIds(contestIds, status)
	var list []interface{}
	for _, v := range contestUserMap {
		list = append(list, v.CreatorId)
	}
	return contestUserMap, list
}

func (c *ContestUserMap) GetMap(ids []interface{}, status int) map[uint64][]*User {
	contestUserMap := c.GetByContestIds(ids, status)
	ret := make(map[uint64][]*User)
	for _, cMap := range contestUserMap {
		if _, ok := ret[cMap.ContestId]; !ok {
			var uList []*User
			ret[cMap.ContestId] = uList
		}
		u := &User{
			Model: Model{
				ID:         cMap.CreatorId,
				Creator:    cMap.Creator,
			},
		}
		ret[cMap.ContestId] = append(ret[cMap.ContestId], u)
	}
	return ret
}

func (c *ContestUserMap) CheckUserSignIn() bool {
	stmt, err := DB.Prepare("SELECT count(1) as cnt FROM contest_user_map WHERE contest_id=? AND creator_id=? AND status=?")
	if err != nil {
		log.Printf("Check User Signin contest error[%v]", err)
		return false
	}
	defer stmt.Close()
	row := stmt.QueryRow(
		&c.ContestId,
		&c.CreatorId,
		STATUS_ENABLE,
	)
	var cnt uint64
	row.Scan(
		&cnt,
	)
	return cnt >= 1
}

func (c *ContestUserMap) CheckUserExists() bool {
	stmt, err := DB.Prepare("SELECT count(1) as cnt FROM contest_user_map WHERE contest_id=? AND creator_id=?")
	if err != nil {
		log.Printf("Check User Exists contest error[%v]", err)
		return false
	}
	defer stmt.Close()
	row := stmt.QueryRow(
		&c.ContestId,
		&c.CreatorId,
	)
	var cnt uint64
	row.Scan(
		&cnt,
	)
	return cnt >= 1
}

func (c *ContestUserMap) ModifyStatus(userIds []interface{}, status int) bool {
	_, err := DB.Exec("UPDATE contest_user_map SET status=" + strconv.Itoa(status) + " WHERE contest_id=" + strconv.FormatUint(c.ContestId, 10) + " AND creator_id IN (?" + strings.Repeat(",?", len(userIds) - 1) + ")", userIds...)
	if err != nil {
		log.Printf("modify lab user status error [%#v]", err)
		return false
	}
	return true
}
