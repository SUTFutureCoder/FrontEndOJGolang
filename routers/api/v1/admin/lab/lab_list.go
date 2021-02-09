package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
)

type ListAdminReq struct {
	LabId uint64 `json:"lab_id"`
	models.Pager
}

type LabListWithSummary struct {
	LabInfo models.Lab
	Summary models.SubmitSummary
	TeseCaseCnt int
}

type ListAdminResp struct {
	LabList []LabListWithSummary
	Count int
}

/**
 * 面向管理员的全局实验室列表，包括不可用实验室
 */
func LabListForAdmin(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	var req ListAdminReq
	var resp ListAdminResp
	err := c.BindJSON(&req)


	// 默认值
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}

	if err != nil {
		log.Printf("bind json error while get lab list[%#v]", err)
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}
	labs, err := models.GetLabList(req.Page, req.PageSize, models.STATUS_ALL)
	resp.Count, err = models.GetLabFullCount()

	if err != nil {
		log.Printf("get db list error while get lab list[%#v]", err)
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	var labIds []interface{}
	for _, lab := range labs {
		labIds = append(labIds, lab.ID)
	}

	labSubmitSummary := models.GetLabSubmitSummary(labIds)
	labTestcaseCnt := models.GetLabTestcaseCntByLabIds(labIds)
	// summary
	for _, lab := range labs {
		var labListWithSummary LabListWithSummary
		labListWithSummary.LabInfo = lab
		if s, ok := labSubmitSummary[lab.ID]; ok {
			labListWithSummary.Summary = *s
		}
		if s, ok := labTestcaseCnt[lab.ID]; ok {
			labListWithSummary.TeseCaseCnt = s
		}
		resp.LabList = append(resp.LabList, labListWithSummary)
	}
	appG.RespSucc(resp)
}

func LabSummaryForAdmin(labIdList []uint64) {

}