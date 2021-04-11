package models

import (
	"FrontEndOJGolang/pkg/utils"
	"database/sql"
	"errors"
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

func (c *Contest) GetByIds(contestIds []interface{}, validCheck bool) []*Contest {
	var contests []*Contest
	if len(contestIds) == 0 {
		return contests
	}
	rows := &sql.Rows{}
	var err error
	if validCheck {
		var sqlQuery []interface{}
		sqlQuery = append(append(sqlQuery, contestIds...), STATUS_ENABLE, utils.GetMillTime(), utils.GetMillTime())
		rows, err = DB.Query("SELECT id, contest_name, contest_desc, contest_start_time, contest_end_time, signup_start_time, signup_end_time, status, creator_id, creator, create_time, update_time FROM contest WHERE id IN (?"+strings.Repeat(",?", len(contestIds)-1)+") AND status=? AND contest_start_time <= ? AND contest_end_time >= ?", sqlQuery...)
	} else {
		rows, err = DB.Query("SELECT id, contest_name, contest_desc, contest_start_time, contest_end_time, signup_start_time, signup_end_time, status, creator_id, creator, create_time, update_time FROM contest WHERE id IN (?"+strings.Repeat(",?", len(contestIds)-1)+")", contestIds...)
	}

	if err != nil {
		log.Printf("get contest list by ids error [%v]\n", err)
		return contests
	}
	defer rows.Close()
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

func (c *Contest) Modify() bool {
	stmt, err := DB.Prepare("UPDATE contest SET contest_name=?, contest_desc=?, contest_start_time=?, contest_end_time=?, signup_start_time=?, signup_end_time=?, update_time=? WHERE id=?")
	if err != nil {
		log.Printf("update contest error [%#v]", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		c.ContestName,
		c.ContestDesc,
		c.ContestStartTime,
		c.ContestEndTime,
		c.SignupStartTime,
		c.SignupEndTime,
		utils.GetMillTime(),
		c.ID,
	)
	if err != nil {
		log.Printf("update contest error [%#v]", err)
		return false
	}
	return true
}

func (c *Contest) ModifyStatus(to int) bool {
	stmt, err := DB.Prepare("UPDATE contest SET status=?, update_time=? WHERE id=?")
	if err != nil {
		log.Printf("update contest status error [%#v]", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(to, utils.GetMillTime(), c.ID)
	if err != nil {
		log.Printf("update contest status error[%#v]", err)
		return false
	}
	return true
}

func (c *Contest) CheckParams() error {
	if c.ContestEndTime == 0 || c.SignupEndTime == 0 {
		return errors.New("time can not equals 0")
	}
	if c.ContestStartTime > c.ContestEndTime || c.SignupStartTime > c.SignupEndTime {
		return errors.New("end time must greater than start time")
	}
	if c.SignupEndTime > c.ContestEndTime {
		return errors.New("contest end time must greater than signup end time")
	}
	if c.ContestEndTime < utils.GetMillTime() {
		return errors.New("end time must greater than now")
	}
	return nil
}
