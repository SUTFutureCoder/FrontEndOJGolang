package models

import "log"

// Contest 比赛表
type Contest struct {
	Model
	// ContestName 比赛名称
	ContestName string `json:"contest_name"`
	// ContestDesc 比赛描述
	ContestDesc string `json:"contest_desc"`
	// ContestStartTime 比赛开始时间
	ContestStartTime uint64 `json:"contest_start_time"`
	// ContestEndTime 比赛结束时间
	ContestEndTime uint64 `json:"contest_end_time"`
	// SignupStartTime
	SignupStartTime uint64 `json:"signup_start_time"`
	// SignupEndTime
	SignupEndTime uint64 `json:"signup_end_time"`
}

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