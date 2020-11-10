package lab_testcase

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReqLabTestCaseList struct {
	LabId uint64 `json:"lab_id" from:"lab_id"`
}

func List(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	var req ReqLabTestCaseList
	var err error
	err = c.BindJSON(&req)
	if req.LabId == 0 {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, "param error")
		return
	}

	_, err = app.GetUserFromSession(c.Request)
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.UNAUTHORIZED, nil)
		return
	}

	testcaseIds, err := models.GetLabTestcaseMapByLabId(req.LabId)
	if len(testcaseIds) == 0 {
		appG.Response(http.StatusOK, e.SUCCESS, nil)
		return
	}
	labTestCases, err := models.GetTestcaseByIds(testcaseIds)

	appG.Response(http.StatusOK, e.SUCCESS, labTestCases)
	return
}
