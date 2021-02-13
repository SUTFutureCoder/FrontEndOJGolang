package models

import (
	"database/sql"
	"log"
	"strings"
)

// LabTestcaseMap 实验室、测试用例关联表
type LabTestcaseMap struct {
	Model
	// LabID 实验室id
	LabID uint64 `json:"lab_id"`
	// TestcaseID 测试用例id
	TestcaseID uint64 `json:"testcase_id"`
}

func GetLabTestcaseMapByLabId(labId uint64) ([]interface{}, error) {
	var testcaseIds []interface{}
	stmt, err := DB.Prepare("SELECT testcase_id FROM lab_testcase_map WHERE lab_id = ? AND status = 1")
	rows, err := stmt.Query(
		&labId,
	)
	defer rows.Close()
	for rows.Next() {
		var testcaseId int
		rows.Scan(&testcaseId)
		testcaseIds = append(testcaseIds, testcaseId)
	}
	return testcaseIds, err
}

func (labTestCaseMap *LabTestcaseMap) Insert(tx *sql.Tx) (sql.Result, error) {
	stmt, err := tx.Prepare("INSERT INTO lab_testcase_map (lab_id, testcase_id, creator_id, creator, create_time) VALUES (?,?,?,?,?)")
	defer stmt.Close()
	result, err := stmt.Exec(
		labTestCaseMap.LabID,
		labTestCaseMap.TestcaseID,
		labTestCaseMap.CreatorId,
		labTestCaseMap.Creator,
		labTestCaseMap.CreateTime,
	)
	return result, err

}

func GetLabTestcaseCntByLabIds(labIds []interface{}) map[uint64]int {
	testcaseCntMap := make(map[uint64]int)
	rows, err := DB.Query("SELECT lab_id, count(*) as cnt FROM lab_testcase_map WHERE lab_id IN (?"+strings.Repeat(",?", len(labIds)-1)+")" + " GROUP BY lab_id", labIds...)
	if err != nil {
		log.Printf("get lab testcase cnt from db error [%#v]", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id uint64
			count int
		)
		err = rows.Scan(&id, &count)
		testcaseCntMap[id] = count
	}
	return testcaseCntMap
}

func (labTestCaseMap *LabTestcaseMap)InvalidLabAllTestcases(tx *sql.Tx) {
	stmt, err := tx.Prepare("UPDATE lab_testcase_map SET status=? WHERE lab_id=?")
	if err != nil {
		tx.Rollback()
		return
	}
	defer stmt.Close()
	stmt.Exec(
			labTestCaseMap.Status,
			labTestCaseMap.LabID,
		)
	return
}