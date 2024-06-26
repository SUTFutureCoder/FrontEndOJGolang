package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

type adminReq struct {
	LabId uint64 `json:"lab_id"`
}

func prepareAdmin(appG *app.Gin) adminReq {
	var req adminReq
	err := appG.C.BindJSON(&req)
	if err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, nil)
		appG.C.Abort()
		return req
	}
	return req
}

func DisableLab(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	req := prepareAdmin(&appG)
	if req.LabId == 0 {
		return
	}
	lab := models.Lab{
		Model: models.Model{ID: req.LabId},
	}
	lab.ModifyStatus(models.STATUS_DISABLE)
	appG.RespSucc(nil)
}

func EnableLab(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	req := prepareAdmin(&appG)
	if req.LabId == 0 {
		return
	}
	lab := models.Lab{
		Model: models.Model{ID: req.LabId},
	}
	lab.ModifyStatus(models.STATUS_ENABLE)
	appG.RespSucc(nil)
}

func ConstructingLab(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	req := prepareAdmin(&appG)
	if req.LabId == 0 {
		return
	}
	lab := models.Lab{
		Model: models.Model{ID: req.LabId},
	}
	lab.ModifyStatus(models.STATUS_CONSTRUCTING)
	appG.RespSucc(nil)
}
