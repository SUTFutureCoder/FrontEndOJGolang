package models

import (
	"FrontEndOJGolang/pkg/setting"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

const (
	STATUS_HIDE 		= -3
	STATUS_CONSTRUCTING = -2
	STATUS_ALL          = -1
	STATUS_DISABLE      = 0
	STATUS_ENABLE       = 1
)

type Model struct {
	// ID 自增ID
	ID uint64 `json:"id"`
	// Status 实验室状态
	Status int `json:"status"`
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

func GetByPager(sqlpre string, page Pager, status int) (*sql.Stmt, *sql.Rows, error){
	DefaultPage(&page.Page, &page.PageSize)
	offset := (page.Page - 1) * page.PageSize

	var stmt *sql.Stmt
	var rows *sql.Rows
	var err error
	if status != STATUS_ALL {
		stmt, err = DB.Prepare( sqlpre + " WHERE status=? ORDER BY id desc LIMIT ? OFFSET ? ")
		rows, err = stmt.Query(
			&status,
			&page.PageSize,
			&offset,
		)
	} else {
		stmt, err = DB.Prepare( sqlpre + " ORDER BY id desc LIMIT ? OFFSET ? ")
		rows, err = stmt.Query(
			&page.PageSize,
			&offset,
		)
	}
	return stmt, rows, err
}

func GetCountByStatus(table string, status int) (int, error) {
	var stmt *sql.Stmt
	var err error
	if status != STATUS_ALL {
		stmt, err = DB.Prepare("SELECT count(1) as cnt FROM " + table + " WHERE status=" + strconv.Itoa(status))
	} else {
		stmt, err = DB.Prepare("SELECT count(1) as cnt FROM " + table)
	}

	defer stmt.Close()
	if err != nil {
		log.Printf("get list count error table[%s] err[%v]\n", table, err)
		return 0, err
	}
	var cnt int
	row := stmt.QueryRow()
	err = row.Scan(&cnt)
	return cnt, err
}

func GetListByIdAndStatus(sqlpre string, id uint64, status int) (*sql.Stmt, *sql.Rows, error) {
	var stmt *sql.Stmt
	var rows *sql.Rows
	var err error
	if status != STATUS_ALL {
		stmt, err = DB.Prepare(sqlpre + " WHERE status=? AND id=?")
		rows, err = stmt.Query(
			&status,
			&id,
		)
	} else {
		stmt, err = DB.Prepare(sqlpre + " WHERE id=?")
		rows, err = stmt.Query(
			&id,
		)
	}
	return stmt, rows, err
}