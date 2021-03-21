package lab_submit

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
)

func SubmitList(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	// get user info from session
	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		return
	}

	pager := models.ToPager(c)
	labSubmit := &models.LabSubmit{}
	labSubmits, err := labSubmit.GetUserSubmits(userSession.Id, pager)
	if err != nil {
		log.Printf("[ERROR] get user lab submits err[%v] userId[%d]", err, userSession.Id)
		appG.RespErr(e.ERROR, nil)
		return
	}
	appG.RespSucc(labSubmits)
}

type submitlistByLabIdReq struct {
	LabId uint64 `json:"lab_id"`
}

func SubmitListByLabId(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	// get user info from session
	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		return
	}

	var req submitlistByLabIdReq
	err := c.BindJSON(&req)
	if err != nil || req.LabId == 0 {
		appG.RespErr(e.PARSE_PARAM_ERROR, nil)
		return
	}
	labSubmit := &models.LabSubmit{}
	labSubmits, err := labSubmit.GetUserSubmitsByLabId(userSession.Id, req.LabId)
	if err != nil {
		log.Printf("[ERROR] get user lab submits by lab ids err[%v] userId[%d] labId[%s]", err, userSession.Id, req.LabId)
		appG.RespErr(e.ERROR, nil)
		return
	}
	appG.RespSucc(labSubmits)
}


type daySubmitsReq struct {
	Time uint64 `json:"time"`
}
type daySubmitsResp struct {
	LabNameHash map[uint64]string `json:"lab_name_hash"`
	Submits []models.LabSubmit `json:"submits"`
}
func DaySubmits(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	var req daySubmitsReq
	var resp daySubmitsResp
	if err := c.BindJSON(&req); err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, nil)
	}
	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		return
	}
	labSubmit := &models.LabSubmit{}
	resp.Submits = labSubmit.GetUserDaySubmits(userSession.Id, req.Time)
	// getLabIds
	var labIds []interface{}
	for _, v := range resp.Submits {
		labIds = append(labIds, v.LabID)
	}

	lab := &models.Lab{}
	labList := lab.GetByIds(labIds)
	// parsehash
	resp.LabNameHash = make(map[uint64]string)
	for _, v := range labList {
		resp.LabNameHash[v.ID] = v.LabName
	}
	appG.RespSucc(resp)
}