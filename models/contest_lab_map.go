package models

import (
	"log"
)

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

func (c *ContestLabMap) GetLabList(pager Pager, status int) ([]*ContestLabMap, error) {
	var contestLabMaps []*ContestLabMap
	stmt, rows, err := GetByPager("SELECT id, contest_id, lab_id, creator_id, status, creator_id, creator, create_time, update_time FROM contest_lab_map", pager, status)
	defer stmt.Close()
	if err != nil {
		log.Printf("get contest lab list from db error [%v]", err)
		return nil, err
	}

	if rows == nil {
		return nil, err
	}
	for rows.Next() {
		c := &ContestLabMap{}
		err = rows.Scan(
			&c.ID, &c.ContestId, &c.LabId, &c.Status, &c.CreatorId, &c.Creator, &c.CreateTime, &c.UpdateTime,
		)
		contestLabMaps = append(contestLabMaps, c)
	}
	return contestLabMaps, err
}
