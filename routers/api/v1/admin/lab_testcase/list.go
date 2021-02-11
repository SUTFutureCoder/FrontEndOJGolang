package lab_testcase

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

type reqLabTestCaseList struct {
	LabId uint64 `json:"lab_id" from:"lab_id"`
}

func List(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	var req reqLabTestCaseList
	err := c.BindJSON(&req)
	if err != nil || req.LabId == 0 {
		appG.RespErr(e.INVALID_PARAMS, "param error")
		return
	}

	testcaseIds, err := models.GetLabTestcaseMapByLabId(req.LabId)
	if len(testcaseIds) == 0 {
		appG.RespSucc(nil)
		return
	}
	labTestCases, err := models.GetTestcaseByIds(testcaseIds)

	appG.RespSucc(labTestCases)
	return
}
