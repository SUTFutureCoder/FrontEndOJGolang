package lab_testcase

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/setting"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type testcaseRunReq struct {
	LabId       uint64             `json:"lab_id"`
	LabTestcase models.LabTestcase `json:"lab_testcase"`
}

type httpTestResult struct {
	Code int
	Msg  string
	Data testResult
}
type testResult struct {
	Id             uint64
	TestCaseId     int
	TestCaseInput  string
	SubmitOutput   string
	TestcaseOutput string
	Status         int
	Err            string
}

func Run(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	var req testcaseRunReq
	var res httpTestResult
	err := c.BindJSON(&req)
	if err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, err.Error())
		return
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		appG.RespErr(e.ERROR, err.Error())
		return
	}
	resp, err := http.Post(fmt.Sprintf("%s:%s/%s", setting.JudgerSetting.JudgerAddr, setting.JudgerSetting.HttpPort, "httpjudger"), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		appG.RespErr(e.ERROR, err.Error())
		return
	}
	json.NewDecoder(resp.Body).Decode(&res)
	if res.Code != e.SUCCESS {
		appG.RespErr(res.Code, res.Data)
		return
	}
	appG.RespSucc(res.Data)
}
