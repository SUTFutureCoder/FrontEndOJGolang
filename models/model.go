package models

import (
	"FrontEndOJGolang/pkg/setting"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	STATUS_CONSTRUCTING = -2
	STATUS_ALL          = -1
	STATUS_DISABLE      = 0
	STATUS_ENABLE       = 1
)

type Model struct {
	// ID 自增ID
	ID uint64 `json:"id"`
	// Status 实验室状态
	Status int8 `json:"status"`
	// CreatorId 创建人Id
	CreatorId uint64 `json:"creator_id"`
	// Creator 创建人
	Creator string `json:"creator"`
	// CreateTime 创建时间
	CreateTime int64 `json:"create_time"`
	// UpdateTime 修改时间
	UpdateTime int `json:"update_time"`
}

type Pager struct {
	// 页数
	Page int `json:"page"`
	// 页面大小
	PageSize int `json:"page_size"`
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

func ToPager(c *gin.Context) Pager {
	var pager Pager
	err := c.BindJSON(&pager)
	if err != nil || pager.Page == 0 || pager.PageSize == 0 {
		pager = Pager{
			Page:     1,
			PageSize: 20,
		}
	}
	return pager
}

func DefaultPage(page, pageSize *int) {
	// 处理默认
	if *page <= 0 {
		*page = 1
	}
	if *pageSize <= 0 {
		*pageSize = 20
	}
}
