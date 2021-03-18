package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
)

type listWithSummaryReq struct {
	LabId uint64 `json:"lab_id"`
	models.Pager
}

type labListWithSummary struct {
	LabInfo     models.Lab           `json:"lab_info"`
	Summary     models.SubmitSummary `json:"summary"`
	TeseCaseCnt int                  `json:"testcase_count"`
}

type listWithSummaryResp struct {
	LabList []labListWithSummary `json:"lab_list"`
	Count   int                  `json:"count"`
}

/**
 * 实验室列表
 * 面向未登录用户及非管理员只显示可用实验室列表
 * 面向管理员的全局实验室列表，包括不可用实验室
 */
func LabListAndSummary(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	var req listWithSummaryReq
	var resp listWithSummaryResp
	err := c.BindJSON(&req)

	userSession := app.GetUserFromSessionNoRespErr(appG)

	if err != nil {
		log.Printf("bind json error while get lab list[%#v]", err)
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	var labs []models.Lab
	status := models.STATUS_ALL
	// guest or no admin mode
	if userSession.Id == 0 || userSession.UserType != models.USERTYPE_ADMIN {
		status = models.STATUS_ENABLE
	}
	lab := &models.Lab{}
	if req.LabId != 0 {
		labs, err = lab.GetListById(req.LabId, status)
		resp.Count = len(labs)
	} else {
		labs, err = lab.GetList(req.Pager, status)
		resp.Count, err = models.GetCountByStatus(models.TABLE_LAB, models.STATUS_ALL)
	}

	if err != nil {
		log.Printf("get db list error while get lab list[%#v]", err)
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	var labIds []interface{}
	for _, lab := range labs {
		labIds = append(labIds, lab.ID)
	}

	labSubmit := &models.LabSubmit{}
	labTestCaseMap := &models.LabTestcaseMap{}
	labSubmitSummary := labSubmit.GetSummary(labIds)
	labTestcaseCnt := labTestCaseMap.GetCntByLabIds(labIds)
	// summary
	for _, lab := range labs {
		var tmpLabListwithsummary labListWithSummary
		tmpLabListwithsummary.LabInfo = lab
		if s, ok := labSubmitSummary[lab.ID]; ok {
			tmpLabListwithsummary.Summary = *s
		}
		if s, ok := labTestcaseCnt[lab.ID]; ok {
			tmpLabListwithsummary.TeseCaseCnt = s
		}
		resp.LabList = append(resp.LabList, tmpLabListwithsummary)
	}
	appG.RespSucc(resp)
}
