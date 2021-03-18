package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

type respLabInfo struct {
	LabInfo models.Lab `json:"lab_info"`
}

type reqLabInfo struct {
	Id uint64 `json:"id" from:"id"`
}

func LabInfo(c *gin.Context) {
	appGin := app.Gin{
		C: c,
	}

	userSession := app.GetUserFromSession(appGin)
	if userSession.Id == 0 {
		return
	}

	var resp respLabInfo
	var req reqLabInfo
	err := c.BindJSON(&req)
	if err != nil {
		appGin.RespErr(e.INVALID_PARAMS, err)
		return
	}
	resp.LabInfo.ID = req.Id
	lab := &models.Lab{}
	err = lab.GetFullInfo(resp.LabInfo.ID)
	if err != nil {
		appGin.RespErr(e.ERROR, err)
		return
	}
	resp.LabInfo = *lab
	if userSession.UserType != models.USERTYPE_ADMIN {
		resp.LabInfo.LabSample = ""
	}

	appGin.RespSucc(resp)
}
