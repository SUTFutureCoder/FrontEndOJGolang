package contest

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

type modifyContest struct {
	ContestId uint64 `json:"contest_id"`
	Status int `json:"status"`
}

func ModifyContestStatus(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	modifyContest := &modifyContest{}
	err := c.BindJSON(modifyContest)
	if err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, nil)
		return
	}
	contest := &models.Contest{
		Model: models.Model{
			ID: modifyContest.ContestId,
		},
	}
	if contest.ID == 0 {
		appG.RespErr(e.INVALID_PARAMS, "contest id equals 0")
		return
	}

	if !contest.ModifyStatus(modifyContest.Status) {
		appG.RespErr(e.ERROR, "modify contest status error")
		return
	}

}
