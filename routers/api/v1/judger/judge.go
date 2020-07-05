package judger

import (
	"FrontEndOJGolang/caroline"
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Judge(c *gin.Context) {

	appG := app.Gin{
		C: c,
	}

	submitIdStr, _ := c.GetPostForm("submit_id")
	submitId, err := strconv.Atoi(submitIdStr)

	// 获取lab_id
	labSubmit, err := models.GetSubmitById(submitId)
	if err != nil {
		log.Printf("")
	}

	if labSubmit == nil || labSubmit.LabID == 0 {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, "实验室id未找到")
		return
	}

	// 获取case信息
	testcaseIds, err := models.GetLabTestcaseMapByLabId(labSubmit.LabID)
	if len(testcaseIds) == 0 {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, "实验室测试列表为空")
		return
	}

	// 获取testcase详情
	testcases, err := models.GetTestcaseByIds(testcaseIds)

	// 执行测试用例
	testChamberFileName := caroline.WriteSubmitToFile(labSubmit)
	caroline.ExecCaroline("file://"+testChamberFileName, testcases)

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}
