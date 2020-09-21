package lab_submit

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

	userSession, err := app.GetUserFromSession(c.Request)
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.UNAUTHORIZED, nil)
		return
	}

	if app.LimitUserSubmitFluency(userSession.Id) {
		appG.Response(http.StatusTooManyRequests, e.TOO_MANY_REQUESTS, nil)
		return
	}

	labSubmit.CreatorId, labSubmit.Creator = userSession.Id, userSession.Name
	labSubmit.CreateTime = time.Now().UnixNano() / 1e6
	lastId, err := labSubmit.Insert()
	if err != nil {
		log.Printf("[ERROR] lab submit error[%v]\n", err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, lastId)

}
