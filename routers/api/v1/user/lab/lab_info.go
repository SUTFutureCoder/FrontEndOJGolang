package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

type RespLabInfo struct {
	LabInfo models.Lab
}

type ReqLabInfo struct {
	Id uint64 `json:"id" from:"id"`
}

func LabInfo(c *gin.Context) {
	appGin := app.Gin{
		C: c,
	}
	var resp RespLabInfo
	var req ReqLabInfo
	var err error
	err = c.BindJSON(&req)
	if err != nil {
		appGin.RespErr(e.INVALID_PARAMS, err)
		return
	}
	resp.LabInfo.ID = req.Id
	resp.LabInfo, err = models.GetLabFullInfo(resp.LabInfo.ID)
	if err != nil {
		appGin.RespErr(e.ERROR, err)
		return
	}
	appGin.RespSucc(resp)
}
