package models

import (
	"FrontEndOJGolang/pkg/utils"
	"database/sql"
	"log"
	"strconv"
	"strings"
)

// Lab 实验室表
type Lab struct {
	Model
	// LabName 实验室名称
	LabName string `json:"lab_name"`
	// LabDesc 实验室描述
	LabDesc string `json:"lab_desc"`
	// LabType 实验室类型
	LabType int8 `json:"lab_type"`
	// LabSample 实验室样例或地址
	LabSample string `json:"lab_sample"`
	// LabTemplate 实验室模板代码
	LabTemplate string `json:"lab_template"`
}

const (
	LABTYPE_NORMAL = iota
	LABTYPE_IMITATE
)

const TABLE_LAB = "lab"

func (lab *Lab) Insert() (int64, error) {
	stmt, err := DB.Prepare("INSERT INTO lab (lab_name, lab_desc, lab_type, lab_sample, lab_template, status, creator_id, creator, create_time) VALUES(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Printf("[ERROR] database exec error input[%v] err[%v]", lab, err)
		return 0, err
	}
	defer stmt.Close()
	ret, err := stmt.Exec(
		lab.LabName,
		lab.LabDesc,
		lab.LabType,
		lab.LabSample,
		lab.LabTemplate,
		lab.Status,
		lab.CreatorId,
		lab.Creator,
		lab.CreateTime,
	)
	if err != nil || ret == nil {
		return 0, err
	}
	return ret.LastInsertId()
}

func (lab *Lab) GetListById(labId uint64, status int) ([]Lab, error) {
	stmt, rows, err := GetListByIdAndStatus("SELECT id, lab_name, lab_type, status, creator_id, creator, create_time, update_time FROM lab", labId, status)
	defer stmt.Close()
	if err != nil {
		log.Printf("get lab list from db error [%v]", err)
		return nil, err
	}

	if rows == nil {
		return nil, err
	}
	var labList []Lab
	for rows.Next() {
		var lab Lab
		err = rows.Scan(
			&lab.ID, &lab.LabName, &lab.LabType, &lab.Status, &lab.CreatorId, &lab.Creator, &lab.CreateTime, &lab.UpdateTime,
		)
		labList = append(labList, lab)
	}
	return labList, err
}

func (lab *Lab) GetList(page Pager, status int) ([]Lab, error) {
	stmt, rows, err := GetByPager("SELECT id, lab_name, lab_type, status, creator_id, creator, create_time, update_time FROM lab", page, status)
	defer stmt.Close()
	if err != nil {
		log.Printf("get lab list from db error [%v]", err)
		return nil, err
	}

	if rows == nil {
		return nil, err
	}
	var labList []Lab
	for rows.Next() {
		var lab Lab
		err = rows.Scan(
			&lab.ID, &lab.LabName, &lab.LabType, &lab.Status, &lab.CreatorId, &lab.Creator, &lab.CreateTime, &lab.UpdateTime,
		)
		labList = append(labList, lab)
	}
	return labList, err
}

func (lab *Lab) GetFullInfo(id uint64) error {
	stmt, err := DB.Prepare("SELECT id, lab_name, lab_desc, lab_type, lab_sample, lab_template, status, creator_id, creator, create_time, update_time FROM lab WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRow(&id)
	err = row.Scan(
		&lab.ID, &lab.LabName, &lab.LabDesc, &lab.LabType, &lab.LabSample, &lab.LabTemplate, &lab.Status, &lab.CreatorId, &lab.Creator, &lab.CreateTime, &lab.UpdateTime,
	)
	if err != nil {
		log.Printf("get lab info error [%v]\n", err)
		return err
	}
	return err
}

func (lab *Lab) ModifyStatus(status int) bool {
	stmt, err := DB.Prepare("UPDATE lab SET status=?, update_time=? WHERE id=?")
	if err != nil {
		log.Printf("update lab status error [%#v]", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(status, utils.GetMillTime(), lab.ID)
	if err != nil {
		log.Printf("update modify status error[%#v]", err)
		return false
	}
	return true
}
func (lab *Lab) Modify() {
	stmt, err := DB.Prepare("UPDATE lab SET lab_name=?, lab_desc=?, lab_type=?, lab_sample=?, lab_template=?, update_time=? WHERE id=?")
	if err != nil {
		log.Printf("update lab status error [%#v]", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(lab.LabName, lab.LabDesc, lab.LabType, lab.LabSample, lab.LabTemplate, utils.GetMillTime(), lab.ID)
	if err != nil {
		log.Printf("update modify status error[%#v]", err)
		return
	}
	return
}

// 无视status直接返回根据id检索内容
func (lab *Lab) GetByIds(labIds []interface{}) []Lab {
	var labs []Lab
	if len(labIds) == 0 {
		return labs
	}
	rows, err := DB.Query("SELECT id, lab_name, lab_desc, lab_type, lab_sample, lab_template, status, creator_id, creator, create_time, update_time FROM lab WHERE id IN (?"+strings.Repeat(",?", len(labIds)-1)+")", labIds...)
	defer rows.Close()
	if err != nil {
		log.Printf("get lab list by ids error [%v]\n", err)
		return labs
	}
	for rows.Next() {
		var lab Lab
		err = rows.Scan(
			&lab.ID, &lab.LabName, &lab.LabDesc, &lab.LabType, &lab.LabSample, &lab.LabTemplate, &lab.Status, &lab.CreatorId, &lab.Creator, &lab.CreateTime, &lab.UpdateTime,
			)
		if err != nil {
			log.Printf("scan lab list by ids ")
			return labs
		}
		labs = append(labs, lab)
	}
	return labs
}

func (lab *Lab) HideLabs(labIds []interface{}, tx *sql.Tx) bool {
	_, err := tx.Exec("UPDATE lab SET status=" + strconv.Itoa(STATUS_HIDE) + " WHERE id IN(?" + strings.Repeat(",?", len(labIds) - 1) + ")", labIds...)
	if err != nil {
		tx.Rollback()
		return false
	}
	return true
}