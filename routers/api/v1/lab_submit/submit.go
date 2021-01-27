package lab_submit

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func Submit(c *gin.Context) {

	appG := app.Gin{
		C: c,
	}

	labIdStr, _ := c.GetPostForm("lab_id")

	labSubmit := models.LabSubmit{}

	labSubmit.LabID, _ = strconv.ParseUint(labIdStr, 10, 64)
	labSubmit.SubmitData, _ = c.GetPostForm("submit_data")

	userSession, err := app.GetUserFromSession(c)
	if err != nil || userSession.Id == 0 {
		appG.RespErr(e.NOT_LOGINED, nil)
		return
	}

	if app.LimitUserSubmitFluency(userSession.Id) {
		appG.RespErr(e.TOO_MANY_REQUESTS, nil)
		return
	}

	labSubmit.CreatorId, labSubmit.Creator = userSession.Id, userSession.Name
	labSubmit.CreateTime = time.Now().UnixNano() / 1e6
	lastId, err := labSubmit.Insert()
	if err != nil {
		log.Printf("[ERROR] lab submit error[%v]\n", err)
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	appG.RespSucc(lastId)

}
