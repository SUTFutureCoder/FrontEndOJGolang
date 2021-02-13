package lab_testcase

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/utils"
	"github.com/gin-gonic/gin"
)

type setTestcaseReq struct {
	LabId	  uint64 `json:"lab_id"`
	Testcases []models.LabTestcase `json:"testcases"`
}

func Set(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {return}

	var req setTestcaseReq
	err := c.BindJSON(&req)
	if err != nil {
		appG.RespErr(e.INVALID_PARAMS, err.Error())
		return
	}

	var formalLabTestcaseMap models.LabTestcaseMap
	formalLabTestcaseMap.LabID = req.LabId
	formalLabTestcaseMap.Status = models.STATUS_DISABLE

	// set trans
	tx, err := models.DB.Begin()
	// invalid all formal testcases
	formalLabTestcaseMap.InvalidLabAllTestcases(tx)
	// set all new testcases
	var labTestCaseMap models.LabTestcaseMap
	labTestCaseMap.LabID = req.LabId
	labTestCaseMap.CreatorId = userSession.Id
	labTestCaseMap.Creator = userSession.Name
	labTestCaseMap.Status = models.STATUS_ENABLE

	for _, v := range req.Testcases {
		v.CreateTime = utils.GetMillTime()
		v.CreatorId = userSession.Id
		v.Creator = userSession.Name
		v.Status = models.STATUS_ENABLE

		labTestCaseMap.CreateTime = utils.GetMillTime()
		labTestCaseMap.TestcaseID, err = v.Insert(tx)
		_, err = labTestCaseMap.Insert(tx)
		if err != nil {
			tx.Rollback()
			break
		}
	}
	if err != nil {
		tx.Rollback()
		appG.RespErr(e.ERROR, err.Error())
		return
	}
	tx.Commit()
	appG.RespSucc(true)
}
