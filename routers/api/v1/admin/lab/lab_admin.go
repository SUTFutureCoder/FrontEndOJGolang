package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

type AdminReq struct {
	LabId uint64 `json:"lab_id"`
}

func prepareAdmin(appG *app.Gin) AdminReq {
	var adminReq AdminReq
	err := appG.C.BindJSON(&adminReq)
	if err != nil {
		appG.RespErr(e.INVALID_PARAMS, nil)
		appG.C.Abort()
		return adminReq
	}
	return adminReq
}

func DisableLab(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	adminReq := prepareAdmin(&appG)
	if adminReq.LabId == 0 {return}
	models.ModifyStatus(adminReq.LabId, models.STATUS_DISABLE)
	appG.RespSucc(nil)
}

func EnableLab(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	adminReq := prepareAdmin(&appG)
	if adminReq.LabId == 0 {return}
	models.ModifyStatus(adminReq.LabId, models.STATUS_ENABLE)
	appG.RespSucc(nil)
}
