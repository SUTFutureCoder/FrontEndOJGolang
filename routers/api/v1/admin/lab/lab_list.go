package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
)

type listAdminReq struct {
	LabId uint64 `json:"lab_id"`
	Pager models.Pager
}

type labListWithSummary struct {
	LabInfo models.Lab
	Summary models.SubmitSummary
	TeseCaseCnt int
}

type listAdminResp struct {
	LabList []labListWithSummary
	Count int
}

/**
 * 面向管理员的全局实验室列表，包括不可用实验室
 */
func LabListForAdmin(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	var req listAdminReq
	var resp listAdminResp
	err := c.BindJSON(&req)

	if err != nil {
		log.Printf("bind json error while get lab list[%#v]", err)
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	labs, err := models.GetLabList(req.Pager, models.STATUS_ALL)
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

func LabSummaryForAdmin(labIdList []uint64) {

}