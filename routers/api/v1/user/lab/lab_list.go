package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

type LabListResp struct {
	LabList []models.Lab
	Count   int
}

type LabListReq struct {
	models.Pager
}

func LabList(c *gin.Context) {
	appGin := app.Gin{
		C: c,
	}

	var req LabListReq
	err := c.BindJSON(&req)
	if err != nil {
		appGin.RespErr(e.INVALID_PARAMS, err)
		return
	}

	var resp LabListResp
	resp.LabList, err = models.GetLabList(req.Page, req.PageSize, models.STATUS_ENABLE)
	resp.Count, err = models.GetLabListCount(models.STATUS_ENABLE)
	if err != nil {
		appGin.RespErr(e.ERROR, err)
		return
	}
	appGin.RespSucc(resp)
}
