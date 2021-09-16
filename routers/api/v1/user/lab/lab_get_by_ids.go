package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

type labGetByIdsReq struct {
	LabIds []uint64 `json:"lab_ids"`
}

type labGetByIdsResp struct {
	LabList []models.Lab `json:"lab_list"`
	Count   int          `json:"count"`
}

func LabGetByIds(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	req := &labGetByIdsReq{}
	resp := &labGetByIdsResp{}
	err := c.BindJSON(req)
	if err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, err)
		return
	}
	if len(req.LabIds) == 0 {
		appG.RespSucc(resp)
		return
	}

	lab := &models.Lab{}
	resp.LabList = lab.GetByIds(models.ConvertUint64ToInterface(req.LabIds))
	for i := range resp.LabList {
		// 脱敏
		resp.LabList[i].HideLabDetail()
	}
	resp.Count = len(resp.LabList)
	appG.RespSucc(resp)
}
