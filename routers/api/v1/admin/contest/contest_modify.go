package contest

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

type contestModifyReq struct {
	models.Contest
}

func ModifyContest(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	req := &contestModifyReq{}
	err := c.BindJSON(req)
	if err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, nil)
		return
	}

	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		return
	}

	contest := &models.Contest{}
	err = contest.CheckParams()
	if err != nil {
		appG.RespErr(e.INVALID_PARAMS, err)
		return
	}

	if !contest.Modify() {
		appG.RespErr(e.ERROR, "modify contest error")
		return
	}
	appG.RespSucc(nil)
}
