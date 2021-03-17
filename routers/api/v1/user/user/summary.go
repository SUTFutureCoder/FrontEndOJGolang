package user

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"github.com/gin-gonic/gin"
)

/**
 * 聚合统计结果
 */
func Summary(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		return
	}

	userSubmitSummary := &models.UserSubmitSummary{}

	userIds := make([]interface{}, 0)
	userIds = append(userIds, userSession.Id)

	labSubmit := &models.LabSubmit{}
	userSubmitsSummary := labSubmit.SummaryUserSubmits(userIds)
	if _, ok := userSubmitsSummary[userSession.Id]; !ok {
		appG.RespSucc(*userSubmitSummary)
		return
	}
	appG.RespSucc(*userSubmitsSummary[userSession.Id])
}

type yearSubmitSummaryResp struct {
	Summary []models.SummaryUserYearSubmit `json:"summary_list"`
}

func YearSubmitSummary(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		return
	}
	var userIds []interface{}
	userIds = append(userIds, userSession.Id)

	labSubmit := &models.LabSubmit{}
	submitSummary := labSubmit.SummaryUserYearSummary(userIds)
	if _, ok := submitSummary[userSession.Id]; !ok {
		appG.RespSucc(nil)
		return
	}

	var resp yearSubmitSummaryResp
	resp.Summary = submitSummary[userSession.Id]
	appG.RespSucc(resp)
}