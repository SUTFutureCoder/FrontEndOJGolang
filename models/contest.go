package models

import (
	"log"
	"strings"
)

// Contest 比赛表
type Contest struct {
	Model
	// ContestName 比赛名称
	ContestName string `json:"contest_name"`
	// ContestDesc 比赛描述
	ContestDesc string `json:"contest_desc"`
	// ContestStartTime 比赛开始时间
	ContestStartTime int64 `json:"contest_start_time"`
	// ContestEndTime 比赛结束时间
	ContestEndTime int64 `json:"contest_end_time"`
	// SignupStartTime
	SignupStartTime int64 `json:"signup_start_time"`
	// SignupEndTime
	SignupEndTime int64 `json:"signup_end_time"`
}

const TABLE_CONTEST = "contest"

func (c *Contest) Insert() (int64, error) {
	stmt, err := DB.Prepare("INSERT INTO contest (contest_name, contest_desc, contest_start_time, contest_end_time, signup_start_time, signup_end_time, status, creator_id, creator, create_time) VALUES(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Printf("[ERROR] database exec error input[%v] err[%v]", c, err)
		return 0, err
	}
	defer stmt.Close()
	ret, err := stmt.Exec(
		c.ContestName,
		c.ContestDesc,
		c.ContestStartTime,
		c.ContestEndTime,
		c.SignupStartTime,
		c.SignupEndTime,
		c.Status,
		c.CreatorId,
		c.Creator,
		c.CreateTime,
	)
	if err != nil || ret == nil {
		return 0, err
	}
	return ret.LastInsertId()
}

func (c *Contest) GetList(page Pager, status int) ([]*Contest, error) {
	var contests []*Contest
	stmt, rows, err := GetByPager("SELECT id, contest_name, contest_desc, contest_start_time, contest_end_time, signup_start_time, signup_end_time, status, creator_id, creator, create_time, update_time FROM contest", page, status)
	defer stmt.Close()
	if err != nil {
		log.Printf("get contest list from db error [%v]", err)
		return nil, err
	}

	if rows == nil {
		return nil, err
	}
	for rows.Next() {
		contest := &Contest{}
		err = rows.Scan(
			&contest.ID, &contest.ContestName, &contest.ContestDesc, &contest.ContestStartTime, &contest.ContestEndTime, &contest.SignupStartTime, &contest.SignupEndTime, &contest.Status, &contest.CreatorId, &contest.Creator, &contest.CreateTime, &contest.UpdateTime,
		)
		contests = append(contests, contest)
	}
	return contests, err
}

func (c *Contest) GetListById(contestId uint64, status int) ([]*Contest, error) {
	stmt, rows, err := GetListByIdAndStatus("SELECT id, contest_name, contest_desc, contest_start_time, contest_end_time, signup_start_time, signup_end_time, status, creator_id, creator, create_time, update_time FROM contest", contestId, status)
	defer stmt.Close()
	if err != nil {
		log.Printf("get contest list from db error [%v]", err)
		return nil, err
	}

	if rows == nil {
		return nil, err
	}
	var contestList []*Contest
	for rows.Next() {
		contest := &Contest{}
		err = rows.Scan(
			&contest.ID, &contest.ContestName, &contest.ContestDesc, &contest.ContestStartTime, &contest.ContestEndTime, &contest.SignupStartTime, &contest.SignupEndTime, &contest.Status, &contest.CreatorId, &contest.Creator, &contest.CreateTime, &contest.UpdateTime,
		)
		contestList = append(contestList, contest)
	}
	return contestList, err
}

func (c *Contest) GetByIds(contestIds []interface{}) []*Contest {
	var contests []*Contest
	if len(contestIds) == 0 {
		return contests
	}
	rows, err := DB.Query("SELECT id, contest_name, contest_desc, contest_start_time, contest_end_time, signup_start_time, sigup_end_time, status, creator_id, creator, create_time, update_time FROM contest WHERE id IN (?"+strings.Repeat(",?", len(contestIds)-1)+")", contestIds...)
	defer rows.Close()
	if err != nil {
		log.Printf("get contest list by ids error [%v]\n", err)
		return contests
	}
	for rows.Next() {
		contest := &Contest{}
		err = rows.Scan(
			&contest.ID, &contest.ContestName, &contest.ContestDesc, &contest.ContestStartTime, &contest.ContestEndTime, &contest.SignupStartTime, &contest.SignupEndTime, &contest.Status, &contest.CreatorId, &contest.Creator, &contest.CreateTime, &contest.UpdateTime,
		)
		if err != nil {
			log.Printf("scan lab list by ids ")
			return contests
		}
		contests = append(contests, contest)
	}
	return contests
}

func (c *Contest) ModifyStatus() {

}
