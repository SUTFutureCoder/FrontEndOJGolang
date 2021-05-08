package contest

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
)

type listWithSummaryReq struct {
	ContestId uint64 `json:"contest_id"`
	Pager models.Pager
}
type listWithSummaryData struct {
	ContestInfo *models.Contest `json:"contest_info"`
	SubmitSummary *models.SubmitSummary `json:"contest_submit_summary"`
	LabCnt int `json:"contest_lab_count"`
	UserCnt int `json:"contest_user_count"`
}
type listWithSummaryResp struct {
	ContestList []listWithSummaryData `json:"contest_list"`
	Count int `json:"count"`
}

func ListWithSummary(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	var req listWithSummaryReq
	err := c.BindJSON(&req)
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

	var resp listWithSummaryResp
	var contests []*models.Contest
	contest := &models.Contest{}
	if req.ContestId != 0 {
		contests, err = contest.GetListById(req.ContestId, status)
		resp.Count = len(contests)
	} else {
		contests, err = contest.GetList(req.Pager, status)
		resp.Count, err = models.GetCountByStatus(models.TABLE_CONTEST, status)
	}
	if err != nil {
		log.Printf("get db list error while get contest list[%#v]", err)
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	// get contestIds
	var contestIds []interface{}
	for _, contest := range contests {
		contestIds = append(contestIds, contest.ID)
	}

	// get labIds by contestIds
	contestLabMap := &models.ContestLabMap{}
	contestLabIdMap, labIds, _ := contestLabMap.GetIdMap(contestIds, status)
	contestUserMap := &models.ContestUserMap{}
	contestUserMaps := contestUserMap.GetMap(contestIds, status)

	// summary
	labSubmit := &models.LabSubmit{}
	labSubmitSummary := labSubmit.GetSummary(labIds)
	resp.ContestList = prepareSummaryRet(contests, contestLabIdMap, contestUserMaps, labSubmitSummary)

	appG.RespSucc(resp)
}

func prepareSummaryRet(contests []*models.Contest, contestLabIdMap map[uint64][]uint64, contestUserMaps map[uint64][]*models.User, labSubmitSummary map[uint64]*models.SubmitSummary) []listWithSummaryData {
	var contestList []listWithSummaryData
	for _, contest := range contests {
		var tmpListWithSummary listWithSummaryData
		tmpListWithSummary.ContestInfo = contest
		tmpListWithSummary.SubmitSummary = &models.SubmitSummary{}
		if _, labMapOk := contestLabIdMap[contest.ID]; labMapOk {
			for _, labId := range contestLabIdMap[contest.ID] {
				if _, summaryMapOk := labSubmitSummary[labId]; summaryMapOk {
					tmpListWithSummary.SubmitSummary.CountSum += labSubmitSummary[labId].CountSum
					tmpListWithSummary.SubmitSummary.CountFail += labSubmitSummary[labId].CountFail
					tmpListWithSummary.SubmitSummary.CountJuding += labSubmitSummary[labId].CountJuding
					tmpListWithSummary.SubmitSummary.CountAc += labSubmitSummary[labId].CountAc
				}
			}
		}

		if _, ok := contestLabIdMap[contest.ID]; ok {
			tmpListWithSummary.LabCnt = len(contestLabIdMap[contest.ID])
		}
		if _, ok := contestUserMaps[contest.ID]; ok {
			tmpListWithSummary.UserCnt = len(contestLabIdMap[contest.ID])
		}
		contestList = append(contestList, tmpListWithSummary)
	}
	return contestList
}