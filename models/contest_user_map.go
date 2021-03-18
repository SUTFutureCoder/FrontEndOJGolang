package models

import (
	"log"
	"strings"
)

// ContestUserMap 比赛用户关联表
type ContestUserMap struct {
	Model
	// ContestId 比赛Id
	ContestId uint64 `json:"contest_id"`
}

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

func (c *ContestUserMap) GetByContestIds(contestIds []interface{}) []*ContestUserMap {
	var contestUserMaps []*ContestUserMap
	if len(contestIds) == 0 {
		return contestUserMaps
	}
	rows, err := DB.Query("SELECT id, contest_id, status, creator_id, creator, create_time, update_time FROM contest_user_map WHERE contest_id IN (?"+strings.Repeat(",?", len(contestIds)-1)+")", contestIds...)
	defer rows.Close()
	if err != nil {
		log.Printf("get contest user list by ids error [%v]\n", err)
		return contestUserMaps
	}
	for rows.Next() {
		c := &ContestUserMap{}
		err = rows.Scan(
			&c.ID, &c.ContestId, &c.Status, &c.CreatorId, &c.Creator, &c.CreateTime, &c.UpdateTime,
		)
		if err != nil {
			log.Printf("scan lab list by ids ")
			return contestUserMaps
		}
		contestUserMaps = append(contestUserMaps, c)
	}
	return contestUserMaps
}


