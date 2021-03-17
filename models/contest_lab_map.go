package models

import "log"

// ContestLabMap 比赛实验室关联表
type ContestLabMap struct {
	Model
	// ContestId 比赛Id
	ContestId uint64 `json:"contest_id"`
	// LabId 实验室Id
	LabId uint64 `json:"lab_id"`
}

func (c *ContestLabMap) Insert() (int64, error) {
	stmt, err := DB.Prepare("INSERT INTO contest_lab_map (contest_id, lab_id, status, creator_id, creator, create_time) VALUES(?,?,?,?,?,?)")
	if err != nil {
		log.Printf("[ERROR] database exec error input[%v] err[%v]", c, err)
		return 0, err
	}
	defer stmt.Close()
	ret, err := stmt.Exec(
		c.ContestId,
		c.LabId,
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


