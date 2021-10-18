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
	LabListIds []interface{} `json:"lab_list_ids"`
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

	appG.RespSucc(infoResp{
		ContestInfo: *contest,
		LabListIds: labIds,
	})
}

type userContestAcLabIdsReq struct {
	ContestId uint64 `json:"contest_id"`
}
type userContestAcLabIdsResp struct {
	AcLabIds []uint64 `json:"ac_lab_ids"`
	ContestLabIds []uint64 `json:"contest_lab_ids"`
	AcNum int `json:"ac_num"`
}
func GetUserContestAcLabIds(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	req := &userContestAcLabIdsReq{}
	err := c.BindJSON(req)
	if err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, err)
		return
	}

	userSession := app.GetUserFromSessionNoRespErr(appG)

	labSubmit := &models.LabSubmit{
		ContestId: req.ContestId,
		Model: models.Model{
			CreatorId: userSession.Id,
		},
	}

	contestLabMap := &models.ContestLabMap{}
	_, labIds, err := contestLabMap.GetIdMap([]interface{}{req.ContestId}, models.STATUS_ENABLE)
	labIdsUint64 := models.ConvertInterfaceToUint64(labIds)
	labIdsUint64Map := make(map[uint64]uint64)
	for _, v := range labIdsUint64 {
		labIdsUint64Map[v] = v
	}

	resp := &userContestAcLabIdsResp{}
	resp.AcLabIds = labSubmit.GetUserContestAcLabIds()
	resp.ContestLabIds = labIdsUint64
	resp.AcNum = 0
	for _, v := range resp.AcLabIds {
		if _, ok := labIdsUint64Map[v]; ok {
			resp.AcNum++
		}
	}

	appG.RespSucc(resp)
}


