package contest

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type contestCreateReq struct {
	models.Contest
}

func CreateContest(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}


	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		return
	}
	req := &contestCreateReq{}
	err := c.BindJSON(req)
	if err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, nil)
		return
	}


	contest := &models.Contest{
		Model: models.Model{
			Status: models.STATUS_ENABLE,
			CreatorId: userSession.Id,
			Creator: userSession.Name,
			CreateTime: utils.GetMillTime(),
		},
		ContestName: req.ContestName,
		ContestDesc: req.ContestDesc,
		ContestStartTime: req.ContestStartTime,
		ContestEndTime: req.ContestEndTime,
		SignupStartTime: req.SignupStartTime,
		SignupEndTime: req.SignupEndTime,
	}

	err = contest.CheckParams()
	if err != nil {
		appG.RespErr(e.INVALID_PARAMS, err.Error())
		return
	}

	lastId, err := contest.Insert()
	if lastId == 0 || err != nil {
		log.Printf("insert contest to db error [%v]", err)
		appG.RespErr(e.ERROR, "insert contest to db error")
		return
	}
	appG.RespSucc(lastId)
}