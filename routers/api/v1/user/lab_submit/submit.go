package lab_submit

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type submitReq struct {
	LabId uint64 `json:"lab_id"`
	SubmitData string `json:"submit_data"`
}

func Submit(c *gin.Context) {

	appG := app.Gin{
		C: c,
	}

	var req submitReq
	c.BindJSON(&req)


	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {return}

	if app.LimitUserSubmitFluency(userSession.Id) {
		appG.RespErr(e.TOO_MANY_REQUESTS, nil)
		return
	}

	var labSubmit models.LabSubmit
	labSubmit.LabID = req.LabId
	labSubmit.SubmitData = req.SubmitData
	labSubmit.CreatorId, labSubmit.Creator = userSession.Id, userSession.Name
	labSubmit.CreateTime = utils.GetMillTime()
	lastId, err := labSubmit.Insert()
	if err != nil {
		log.Printf("[ERROR] lab submit error[%v]\n", err)
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	appG.RespSucc(lastId)

}
