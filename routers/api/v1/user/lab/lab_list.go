package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

type labListResp struct {
	LabList []models.Lab `json:"lab_list"`
	Count   int          `json:"count"`
}

type labListReq struct {
	LabId uint64 `json:"lab_id"`
	Pager models.Pager
}

func LabList(c *gin.Context) {
	appGin := app.Gin{
		C: c,
	}

	var req labListReq
	err := c.BindJSON(&req)
	if err != nil {
		appGin.RespErr(e.PARSE_PARAM_ERROR, err)
		return
	}
	app.GetUserFromSession(appGin)
	var resp labListResp
	lab := &models.Lab{}
	if req.LabId != 0 {
		resp.LabList, err = lab.GetListById(req.LabId, models.STATUS_ENABLE)
		resp.Count = len(resp.LabList)
	} else {
		resp.LabList, err = lab.GetList(req.Pager, models.STATUS_ENABLE)
		resp.Count, err = models.GetCountByStatus(models.TABLE_LAB, models.STATUS_ENABLE)
	}
	if err != nil {
		appGin.RespErr(e.ERROR, err)
		return
	}
	appGin.RespSucc(resp)
}
