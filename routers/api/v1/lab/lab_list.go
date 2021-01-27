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
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
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

	// 默认值
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}

	var resp LabListResp
	resp.LabList, err = models.GetLabList(req.Page, req.PageSize)
	resp.Count, err = models.GetLabListCount()
	if err != nil {
		appGin.RespErr(e.ERROR, err)
		return
	}
	appGin.RespSucc(resp)
}
