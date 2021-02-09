package lab_submit

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
)

func SubmitList(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	// get user info from session
	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {return}

	pager := models.ToPager(c)
	labSubmits, err := models.GetUserLabSubmits(userSession.Id, pager)
	if err != nil {
		log.Printf("[ERROR] get user lab submits err[%v] userId[%d]", err, userSession.Id)
		appG.RespErr(e.ERROR, nil)
		return
	}
	appG.RespSucc(labSubmits)
}

type SubmitlistByLabIdReq struct {
	LabId uint64 `json:"lab_id"`
}

func SubmitListByLabId(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	// get user info from session
	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {return}

	var req SubmitlistByLabIdReq
	err := c.BindJSON(&req)
	if err != nil || req.LabId == 0 {
		appG.RespErr(e.INVALID_PARAMS, "invalid params")
		return
	}

	labSubmits, err := models.GetUserLabSubmitsByLabId(userSession.Id, req.LabId)
	if err != nil {
		log.Printf("[ERROR] get user lab submits by lab ids err[%v] userId[%d] labId[%s]", err, userSession.Id, req.LabId)
		appG.RespErr(e.ERROR, nil)
		return
	}
	appG.RespSucc(labSubmits)
}
