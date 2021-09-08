package contest

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

type infoReq struct {
	ContestId uint64 `json:"contest_id"`
}

type infoResp struct {
	ContestInfo models.Contest `json:"contest_info"`
	LabList []models.Lab `json:"lab_list"`
}

func Info(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	req := &infoReq{}
	err := c.BindJSON(req)
	if err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, err)
		return
	}

	userSession := app.GetUserFromSessionNoRespErr(appG)
	status := models.STATUS_ALL
	// guest or no admin mode
	if userSession.Id == 0 || userSession.UserType != models.USERTYPE_ADMIN {
		status = models.STATUS_ENABLE
	}

	// get contest info
	contest := &models.Contest{}
	contests, err := contest.GetListById(req.ContestId, status)

	if len(contests) == 0 {
		appG.RespErr(e.INVALID_PARAMS, "Contest Was Missing")
		return
	}
	contest = contests[0]

	// protect future contest
	if status != models.STATUS_ALL {
		reason := checkAccess(contest, userSession.Id)
		if reason != "" {
			appG.RespErr(e.CONTEST_ACCESS_DENIED, reason)
			return
		}
	}

	contestLabMap := &models.ContestLabMap{}
	_, labIds, err := contestLabMap.GetIdMap([]interface{}{contest.ID}, status)
	labList := make([]models.Lab, 0)

	if len(labIds) != 0 {
		lab := &models.Lab{}
		labList = lab.GetByIds(labIds)
		for k, _ := range labList {
			labList[k].LabSample = ""
		}
	}

	appG.RespSucc(infoResp{
		ContestInfo: *contest,
		LabList: labList,
	})
}



