package contest_lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
)

type manageLabsReq struct {
	ContestId uint64 `json:"contest_id"`
	LabIds []uint64 `json:"lab_ids"`
}

func ManageLabs(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	req := &manageLabsReq{}
	err := c.BindJSON(req)
	if err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, nil)
		return
	}

	if len(req.LabIds) == 0 || req.ContestId == 0 {
		appG.RespErr(e.INVALID_PARAMS, "lab id list or contest id equals 0")
		return
	}

	usersession := app.GetUserFromSession(appG)

	// get lab info
	lab := &models.Lab{}
	labIdsInfo := lab.GetByIds(models.ConvertUint64ToInterface(req.LabIds))
	if len(labIdsInfo) != len(req.LabIds) {
		appG.RespErr(e.INVALID_PARAMS, "get lab ids from db error, some labs not exists")
		return
	}

	// start trans
	// update formal to 0
	// set trans
	tx, err := models.DB.Begin()

	// STEP1 hide lab
	if !lab.HideLabs(models.ConvertUint64ToInterface(req.LabIds), tx) {
		tx.Rollback()
		appG.RespErr(e.ERROR, "hide lab failed")
		return
	}

	// STEP2 invalid former labs
	contestLabMap := &models.ContestLabMap{
		ContestId: req.ContestId,
		Model: models.Model{
			CreatorId: usersession.Id,
			Creator: usersession.Name,
			CreateTime: utils.GetMillTime(),
			Status: models.STATUS_ENABLE,
		},
	}
	if !contestLabMap.InvalidAll(tx) {
		tx.Rollback()
		appG.RespErr(e.ERROR, "invalid all former lab of contest failed")
		return
	}

	// STEP3 insert to contest labs
	if !buildAndInsertBatch(contestLabMap, req.LabIds, tx) {
		tx.Rollback()
		appG.RespErr(e.ERROR, "insert to contest labs error")
		return
	}

	if err != nil {
		tx.Rollback()
		log.Printf("DB set contest lab error [%#v]", err)
		appG.RespErr(e.ERROR, "db error")
		return
	}
	tx.Commit()
	appG.RespSucc(nil)
}

func buildAndInsertBatch(contestLabMap *models.ContestLabMap, labIds []uint64, tx *sql.Tx) bool {
	for _, id := range labIds {
		contestLabMap.LabId = id
		ret, err := contestLabMap.InsertWithTx(tx)
		if err != nil || ret == 0 {
			log.Printf("buildAndInsertBatch failed error[%#v]", err)
			return false
		}
	}
	return true
}
