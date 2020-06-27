package lab_testcase

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Add(c *gin.Context)  {
	appG := app.Gin{
		C: c,
	}

	labId, _ := c.GetPostForm("lab_id")
	testcaseDesc, _ := c.GetPostForm("testcase_desc")
	testcaseCode, _ := c.GetPostForm("testcase_code")
	input, _ := c.GetPostForm("input")
	output, _ := c.GetPostForm("output")
	timeLimitStr, _ := c.GetPostForm("timelimit")
	timeLimit, _ := strconv.Atoi(timeLimitStr)
	memLimitStr, _ := c.GetPostForm("memlimit")
	memLimit, _ := strconv.Atoi(memLimitStr)
	creator := "CaveJohson"
	createTime := time.Now().UnixNano() / 1e6


	tx, err := models.DB.Begin()

	stmt, err := tx.Prepare("INSERT INTO lab_testcase (testcase_desc, testcase_code, input, output, time_limit, mem_limit,  creator, create_time) VALUES (?,?,?,?,?,?,?,?)")
	result, err := stmt.Exec(
		&testcaseDesc,
		&testcaseCode,
		&input,
		&output,
		&timeLimit,
		&memLimit,
		&creator,
		&createTime,
		)

	labTestCaseLastId, err := result.LastInsertId()


	stmt, err = tx.Prepare("INSERT INTO lab_testcase_map (lab_id, testcase_id, creator, create_time) VALUES (?,?,?,?)")
	result, err = stmt.Exec(
			&labId,
			&labTestCaseLastId,
			&creator,
			&createTime,
		)
	defer stmt.Close()

	err = tx.Commit()
	if err != nil {
		log.Printf("[ERROR] add testcase error [%v]\n", err)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}