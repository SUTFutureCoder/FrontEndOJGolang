package contest

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/utils"
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

	contestUserMap := &models.ContestUserMap{
		ContestId: req.ContestId,
		Model : models.Model{
			CreatorId: userSession.Id,
		},
	}
	// protect future contest
	if status != models.STATUS_ALL {
		if contest.ContestStartTime > utils.GetMillTime() {
			appG.RespErr(e.UNAUTHORIZED, "Contest Not Start")
			return
		}
		// check user have signed
		if !contestUserMap.CheckUserSignIn() {
			appG.RespErr(e.UNAUTHORIZED, "Please Signin The Contest First")
			return
		}
	}

	contestLabMap := &models.ContestLabMap{}
	_, labIds, err := contestLabMap.GetIdMap([]interface{}{contest.ID}, status)
	if len(labIds) == 0 {
		appG.RespErr(e.ERROR, "Contest Contains No Labs")
		return
	}

	lab := &models.Lab{}
	labList := lab.GetByIds(labIds)
	for k, _ := range labList {
		labList[k].LabSample = ""
	}
	appG.RespSucc(infoResp{
		ContestInfo: *contest,
		LabList: labList,
	})
}



