package models

import (
	"FrontEndOJGolang/pkg/setting"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Model struct {
	// ID 自增ID
	ID uint64 `json:"id"`
	// Status 实验室状态
	Status int8 `json:"status"`
	// Creator 创建人
	Creator string `json:"creator"`
	// CreateTime 创建时间
	CreateTime int64 `json:"create_time"`
	// UpdateTime 修改时间
	UpdateTime int `json:"update_time"`
}

var DB *sql.DB

func Setup() {
	var err error
	DB, err = sql.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))
	if err != nil {
		log.Fatalf("[FATAL] Database init error[%v]", err)
		return
	}

	DB.SetMaxIdleConns(setting.DatabaseSetting.MaxIdleConns)
	DB.SetMaxOpenConns(setting.DatabaseSetting.MaxOpenConns)
	err = DB.Ping()
	if err != nil {
		log.Fatalf("[FATAL] Database ping error[%v]", err)
	}
}