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

func Add(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	labIdStr, _ := c.GetPostForm("lab_id")
	labId, _ := strconv.ParseInt(labIdStr, 10, 64)
	testcaseDesc, _ := c.GetPostForm("testcase_desc")
	testcaseCode, _ := c.GetPostForm("testcase_code")
	input, _ := c.GetPostForm("input")
	output, _ := c.GetPostForm("output")
	timeLimitStr, _ := c.GetPostForm("timelimit")
	timeLimit, _ := strconv.Atoi(timeLimitStr)
	memLimitStr, _ := c.GetPostForm("memlimit")
	memLimit, _ := strconv.Atoi(memLimitStr)
	waitBeforeStr, _ := c.GetPostForm("wait_before")
	waitBefore, _ := strconv.Atoi(waitBeforeStr)
	creator := "CaveJohson"
	createTime := time.Now().UnixNano() / 1e6

	tx, err := models.DB.Begin()

	labTestCase := models.LabTestcase{
		Model: models.Model{
			Creator:    creator,
			CreateTime: createTime,
		},
		TestcaseDesc: testcaseDesc,
		TestcaseCode: testcaseCode,
		Input:        input,
		Output:       output,
		TimeLimit:    timeLimit,
		MemLimit:     memLimit,
		WaitBefore:   waitBefore,
	}

	labTestCaseLastId, err := models.InsertLabTestCase(tx, &labTestCase)

	labTestCaseMap := models.LabTestcaseMap{
		Model: models.Model{
			Creator:    creator,
			CreateTime: createTime,
		},
		LabID:      labId,
		TestcaseID: labTestCaseLastId,
	}

	_, err = models.InsertLabTestCaseMap(tx, &labTestCaseMap)

	err = tx.Commit()
	if err != nil {
		log.Printf("[ERROR] add testcase error [%v]\n", err)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
