package models

import "log"

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


