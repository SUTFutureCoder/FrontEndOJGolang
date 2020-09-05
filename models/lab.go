package models

import (
	"errors"
	"log"
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
}

const (
	LABTYPE_HTML = iota
	LABTYPE_CSS
	LABTYPE_JS
	LABTYPE_VUE
	LABTYPE_COMPLEX
	LABTYPE_PRD
	LABTYPE_IMITATE
	LABTYPE_SECURITY
	LABTYPE_OTHER
)

func (lab *Lab) Insert() error {
	stmt, err := DB.Prepare("INSERT INTO lab (lab_name, lab_desc, lab_type, lab_sample, creator_id, creator, create_time) VALUES(?,?,?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		log.Printf("[ERROR] database exec error input[%v] err[%v]", lab, err)
		return err
	}
	_, err = stmt.Exec(
		lab.LabName,
		lab.LabDesc,
		lab.LabType,
		lab.LabSample,
		lab.CreatorId,
		lab.Creator,
		lab.CreateTime,
	)
	return nil
}

func GetLabList(page, pageSize int) ([]Lab, error) {
	if pageSize <= 0 || page <= 0 {
		return nil, errors.New("page or pagesize not available")
	}
	stmt, err := DB.Prepare("SELECT id, lab_name, lab_type, creator_id, creator, create_time, update_time FROM lab WHERE status = 1 ORDER BY id desc LIMIT ? OFFSET ? ")
	defer stmt.Close()
	if err != nil {
		log.Printf("get lab list from db error [%v]", err)
		return nil, err
	}
	offset := (page - 1) * pageSize
	rows, err := stmt.Query(
			&pageSize,
			&offset,
		)
	var labList []Lab
	for rows.Next() {
		var lab Lab
		err = rows.Scan(
				&lab.ID, &lab.LabName, &lab.LabType, &lab.CreatorId, &lab.Creator, &lab.CreateTime, &lab.UpdateTime,
			)
		labList = append(labList, lab)
	}
	return labList, err
}

func GetLabListCount() (int, error) {
	stmt, err := DB.Prepare("SELECT count(1) as cnt FROM lab WHERE status = 1")
	defer stmt.Close()
	if err != nil {
		log.Printf("get lab list count error [%v]\n", err)
		return 0, err
	}
	var cnt int
	row := stmt.QueryRow()
	err = row.Scan(&cnt)
	return cnt, err
}

func GetLabInfo(id uint64) (Lab, error) {
	var lab Lab
	stmt, err := DB.Prepare("SELECT id, lab_name, lab_desc, lab_type, lab_sample, status, creator_id, creator, create_time, update_time FROM lab WHERE id=?")
	if err != nil {
		return lab, err
	}
	row := stmt.QueryRow(&id)
	err = row.Scan(
		&lab.ID, &lab.LabName, &lab.LabDesc, &lab.LabType, &lab.LabSample, &lab.Status, &lab.CreatorId, &lab.Creator, &lab.CreateTime, &lab.UpdateTime,
		)
	if err != nil {
		log.Printf("get lab info error [%v]\n", err)
		return lab, err
	}
	return lab, err
}
