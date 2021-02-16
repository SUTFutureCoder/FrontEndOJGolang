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
	userSubmitsSummary := models.SummaryUserSubmits(userIds)
	if _, ok := userSubmitsSummary[userSession.Id]; !ok {
		appG.RespSucc(*userSubmitSummary)
		return
	}
	appG.RespSucc(*userSubmitsSummary[userSession.Id])
}
