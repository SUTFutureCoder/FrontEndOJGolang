package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

type labGetByIdsWithUserSummaryReq struct {
	LabIds []uint64 `json:"lab_ids"`
	ContestId uint64 `json:"contest_id"`
}
type labGetByIdsWithUserSummaryResp struct {
	listWithSummaryResp
}

func LabGetByIdsWithUserSummary(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	req := &labGetByIdsWithUserSummaryReq{}
	resp := &labGetByIdsWithUserSummaryResp{}
	err := c.BindJSON(req)
	if err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, err)
		return
	}
	if len(req.LabIds) == 0 {
		appG.RespSucc(resp)
		return
	}

	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		appG.RespSucc(resp)
		return
	}

	lab := &models.Lab{}
	labList := lab.GetByIds(models.ConvertUint64ToInterface(req.LabIds))
	resp.Count = len(labList)

	// summary
	labSubmit := &models.LabSubmit{}
	labTestCaseMap := &models.LabTestcaseMap{}
	labSubmitSummary := labSubmit.GetSummaryByUserId(models.ConvertUint64ToInterface(req.LabIds), req.ContestId, userSession.Id)
	labTestcaseCnt := labTestCaseMap.GetCntByLabIds(models.ConvertUint64ToInterface(req.LabIds))

	// prepare resp
	for _, lab := range labList {
		var tmpLabListwithsummary labListWithSummary
		tmpLabListwithsummary.LabInfo = lab
		if s, ok := labSubmitSummary[lab.ID]; ok {
			tmpLabListwithsummary.Summary = *s
		}
		if s, ok := labTestcaseCnt[lab.ID]; ok {
			tmpLabListwithsummary.TeseCaseCnt = s
		}
		lab.HideLabDetail()
		resp.LabList = append(resp.LabList, tmpLabListwithsummary)
	}
	appG.RespSucc(resp)
}
