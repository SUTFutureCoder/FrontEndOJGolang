package lab_testcase

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
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
	for _, v := range req.Testcases {
		var labTestCaseMap models.LabTestcaseMap
		labTestCaseMap.TestcaseID, err = v.Insert(tx)
		_, err = labTestCaseMap.Insert(tx)
	}
	tx.Commit()
	appG.RespSucc(nil)
}
