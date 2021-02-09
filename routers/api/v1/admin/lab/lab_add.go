package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type AddLabReq struct {
	LabName     string `json:"lab_name"`
	LabDesc     string `json:"lab_desc"`
	LabSample   string `json:"lab_sample"`
	LabTemplate string `json:"lab_template"`
	LabType     int8   `json:"lab_type"`
}

func AddLab(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {return}

	addLabReq := AddLabReq{}
	err := c.BindJSON(&addLabReq)
	if err != nil {
		appG.RespErr(e.INVALID_PARAMS, err.Error())
		return
	}

	lab := prepareAdd(&addLabReq, &userSession)

	labId, err := lab.Insert()
	if err != nil {
		log.Printf("[ERROR] database exec error input[%v] err[%v]", lab, err)
		appG.RespErr(e.ERROR, nil)
		return
	}

	appG.RespSucc(labId)
}

func prepareAdd(addLabReq *AddLabReq, userSession *app.UserSession) *models.Lab {
	lab := &models.Lab{
		LabName:     addLabReq.LabName,
		LabDesc:     addLabReq.LabDesc,
		LabSample:   addLabReq.LabSample,
		LabTemplate: addLabReq.LabTemplate,
		LabType:     addLabReq.LabType,
	}
	lab.CreatorId, lab.Creator = userSession.Id, userSession.Name
	lab.CreateTime = utils.GetMillTime()
	return lab
}
