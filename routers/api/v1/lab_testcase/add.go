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

	userSession, err := app.GetUserFromSession(c.Request)
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.UNAUTHORIZED, nil)
		return
	}

	labTestCase := models.LabTestcase{}
	labTestCaseMap := models.LabTestcaseMap{}

	prepare(&labTestCase, labTestCaseMap, c, userSession)

	tx, err := models.DB.Begin()
	labTestCaseMap.TestcaseID, err = labTestCase.Insert(tx)
	_, err = labTestCaseMap.Insert(tx)
	err = tx.Commit()
	if err != nil {
		log.Printf("[ERROR] add testcase error [%v]\n", err)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func prepare(labTestCase *models.LabTestcase, labTestCaseMap models.LabTestcaseMap, c *gin.Context, session app.UserSession) {
	labIdStr, _ := c.GetPostForm("lab_id")
	labTestCaseMap.LabID, _ = strconv.ParseInt(labIdStr, 10, 64)
	labTestCase.TestcaseDesc, _ = c.GetPostForm("testcase_desc")
	labTestCase.TestcaseCode, _ = c.GetPostForm("testcase_code")
	labTestCase.Input, _ = c.GetPostForm("input")
	labTestCase.Output, _ = c.GetPostForm("output")
	timeLimitStr, _ := c.GetPostForm("timelimit")
	labTestCase.TimeLimit, _ = strconv.Atoi(timeLimitStr)
	memLimitStr, _ := c.GetPostForm("memlimit")
	labTestCase.MemLimit, _ = strconv.Atoi(memLimitStr)
	waitBeforeStr, _ := c.GetPostForm("wait_before")
	labTestCase.WaitBefore, _ = strconv.Atoi(waitBeforeStr)
	labTestCase.CreatorId, labTestCaseMap.CreatorId = session.Id, session.Id
	labTestCase.Creator, labTestCaseMap.Creator = session.Name, session.Name
	labTestCase.CreateTime = time.Now().UnixNano() / 1e6
}
