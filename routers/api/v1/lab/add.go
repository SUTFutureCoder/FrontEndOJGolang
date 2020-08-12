package lab

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

func AddLab(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	userSession, err := app.GetUserFromSession(c.Request)
	if err != nil {
		log.Printf("[ERROR] get user session error[%v]\n", err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	lab := models.Lab{}

	prepare(lab, c, userSession)

	err = lab.Insert()
	if err != nil {
		log.Printf("[ERROR] database exec error input[%v] err[%v]", lab, err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func prepare(lab models.Lab, c *gin.Context, userSession app.UserSession) {
	lab.LabName, _ = c.GetPostForm("lab_name")
	lab.LabDesc, _ = c.GetPostForm("lab_desc")
	lab.LabSample, _ = c.GetPostForm("lab_sample")
	labTypeStr, _ := c.GetPostForm("lab_type")
	labType, _ := strconv.ParseInt(labTypeStr, 10, 8)
	lab.LabType = int8(labType)
	lab.CreatorId, lab.Creator = userSession.Id, userSession.Name
	lab.CreateTime = time.Now().UnixNano() / 1e6
}
