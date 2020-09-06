package lab_submit

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SubmitList(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	// get user info from session
	userSession, err := app.GetUserFromSession(c.Request)
	if err != nil || userSession.Id == 0 {
		appG.Response(http.StatusUnauthorized, e.UNAUTHORIZED, "please login")
		return
	}

	pager := models.ToPager(c)
	labSubmits, err := models.GetUserLabSubmits(userSession.Id, pager)
	if err != nil {
		log.Printf("[ERROR] get user lab submits err[%v] userId[%d]", err, userSession.Id)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, labSubmits)
}

func SubmitListByLabId(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	// get user info from session
	userSession, err := app.GetUserFromSession(c.Request)
	if err != nil || userSession.Id == 0 {
		appG.Response(http.StatusUnauthorized, e.UNAUTHORIZED, "please login")
		return
	}

	labId, _ := c.GetPostForm("lab_id")
	if labId == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	labSubmits, err := models.GetUserLabSubmitsByLabId(userSession.Id, labId)
	if err != nil {
		log.Printf("[ERROR] get user lab submits by lab ids err[%v] userId[%d] labId[%s]", err, userSession.Id, labId)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, labSubmits)
}
