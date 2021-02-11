package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

type labListResp struct {
	LabList []models.Lab
	Count   int
}

type labListReq struct {
	Pager models.Pager
}

func LabList(c *gin.Context) {
	appGin := app.Gin{
		C: c,
	}

	var req labListReq
	err := c.BindJSON(&req)
	if err != nil {
		appGin.RespErr(e.INVALID_PARAMS, err)
		return
	}

	var resp labListResp
	resp.LabList, err = models.GetLabList(req.Pager, models.STATUS_ENABLE)
	resp.Count, err = models.GetLabListCount(models.STATUS_ENABLE)
	if err != nil {
		appGin.RespErr(e.ERROR, err)
		return
	}
	appGin.RespSucc(resp)
}
