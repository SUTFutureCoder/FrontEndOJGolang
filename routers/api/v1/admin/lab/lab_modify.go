package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

type modifyReq struct {
	Id uint64 `json:"lab_id"`
	// LabName 实验室名称
	LabName string `json:"lab_name"`
	// LabDesc 实验室描述
	LabDesc string `json:"lab_desc"`
	// LabType 实验室类型
	LabType int8 `json:"lab_type"`
	// LabSample 实验室样例或地址
	LabSample string `json:"lab_sample"`
	// LabTemplate 实验室模板代码
	LabTemplate string `json:"lab_template"`
}

func ModifyLab(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	var req modifyReq
	err := c.BindJSON(&req)
	if err != nil || req.Id == 0 {
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	// get lab
	lab := &models.Lab{}
	err = lab.GetLabFullInfo(req.Id)
	if err != nil {
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	// modify lab
	lab.LabName = req.LabName
	lab.LabDesc = req.LabDesc
	lab.LabTemplate = req.LabTemplate
	lab.LabType = req.LabType
	lab.LabSample = req.LabSample

	lab.Modify()

}
