package models

import (
	"database/sql"
	"log"
	"strings"
)

// ContestLabMap 比赛实验室关联表
type ContestLabMap struct {
	Model
	// ContestId 比赛Id
	ContestId uint64 `json:"contest_id"`
	// LabId 实验室Id
	LabId uint64 `json:"lab_id"`
}

func (c *ContestLabMap) InsertWithTx(tx *sql.Tx) (int64, error) {
	stmt, err := tx.Prepare("INSERT INTO contest_lab_map (contest_id, lab_id, status, creator_id, creator, create_time) VALUES(?,?,?,?,?,?)")
	if err != nil {
		tx.Rollback()
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
		tx.Rollback()
		return 0, err
	}
	return ret.LastInsertId()
}
//
//func (c *ContestLabMap) GetLabList(pager Pager, status int) ([]*ContestLabMap, error) {
//	var contestLabMaps []*ContestLabMap
//	stmt, rows, err := GetByPager("SELECT id, contest_id, lab_id, creator_id, status, creator_id, creator, create_time, update_time FROM contest_lab_map", pager, status)
//	defer stmt.Close()
//	if err != nil {
//		log.Printf("get contest lab list from db error [%v]", err)
//		return nil, err
//	}
//
//	if rows == nil {
//		return nil, err
//	}
//	for rows.Next() {
//		c := &ContestLabMap{}
//		err = rows.Scan(
//			&c.ID, &c.ContestId, &c.LabId, &c.Status, &c.CreatorId, &c.Creator, &c.CreateTime, &c.UpdateTime,
//		)
//		contestLabMaps = append(contestLabMaps, c)
//	}
//	return contestLabMaps, err
//}

func (c *ContestLabMap) GetIdMap(contestIds []interface{}, status int) (map[uint64][]uint64, []interface{}, error) {
	contestLabIdMap := make(map[uint64][]uint64)
	var labIdList []interface{}
	if len(contestIds) == 0 {
		return contestLabIdMap, labIdList, nil
	}

	query := "SELECT contest_id, lab_id, status FROM contest_lab_map WHERE contest_id IN (?"+strings.Repeat(",?", len(contestIds)-1)+") ORDER BY lab_order"
	if status != STATUS_ALL {
		contestIds = append(contestIds, status)
		query += " AND status=?"
	}
	rows, err := DB.Query(query, contestIds...)
	defer rows.Close()
	if err != nil {
		log.Printf("get contest lab map ids error [%v]", err)
	}

	for rows.Next() {
		contestLabMap := &ContestLabMap{}
		err = rows.Scan(
			&contestLabMap.ContestId, &contestLabMap.LabId, &contestLabMap.Status,
		)
		if _, ok := contestLabIdMap[contestLabMap.ContestId]; !ok {
			var tmpIdList []uint64
			contestLabIdMap[contestLabMap.ContestId] = tmpIdList
		}
		contestLabIdMap[contestLabMap.ContestId] = append(contestLabIdMap[contestLabMap.ContestId], contestLabMap.LabId)
		labIdList = append(labIdList, contestLabMap.LabId)
	}

	return contestLabIdMap, labIdList, err
}

func (c *ContestLabMap) InvalidAll(tx *sql.Tx) bool {
	if c.ID == 0 {
		tx.Rollback()
		return false
	}
 	stmt, err := tx.Prepare("UPDATE contest_lab_map SET status=? WHERE contest_id=?")
 	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return false
	}
	_, err = stmt.Exec(
			STATUS_DISABLE,
			c.ContestId,
		)
	if err != nil {
		tx.Rollback()
		return false
	}
	return true
}

