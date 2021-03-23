package contest

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type dashboardReq struct {
	ContestId uint64 `json:"contest_id"`
}

type dashboardUserData struct {
	SubmitTimes int `json:"submit_times"`
	TimeSum int `json:"time_sum"`
	IsAc bool `json:"is_ac"`
}

type dashboardResp struct {
	RankDataMap map[uint64]map[uint64]dashboardUserData `json:"rank_data"`
	ContestInfo models.Contest `json:"contest_info"`
	UserInfos []*models.ContestUserMap `json:"user_infos"`
	LabInfos []models.Lab `json:"lab_infos"`
}

// dashboard 仅统计开始到结束时间段的提交分数
func Dashboard(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	req := &dashboardReq{}
	if err := c.BindJSON(req); err != nil {
		appG.RespErr(e.INVALID_PARAMS, "parse dashboard param error")
		return
	}

	contest := &models.Contest{}
	if !checkParams(req, contest, &appG) {
		return
	}

	contestUserMap := &models.ContestUserMap{}
	userInfoList, userIdList := contestUserMap.GetIdListByContestIds([]interface{}{contest.ID}, models.STATUS_ENABLE)

	contestLabMap := &models.ContestLabMap{}
	_, labIdList, err := contestLabMap.GetIdMap([]interface{}{contest.ID}, models.STATUS_ENABLE)
	if err != nil {
		log.Printf("get dashboard contest lab map error [%#v] data [%#v]", err, contest)
		appG.RespErr(e.ERROR, "get dashboard contest lab map error")
		return
	}

	// get lab id
	lab := &models.Lab{}
	labInfoList := lab.GetByIds(labIdList)

	labSubmit := &models.LabSubmit{}
	submitGroupData := labSubmit.GroupByUserAndLabIds(labIdList, userIdList)

	resp := &dashboardResp{}
	resp.UserInfos = userInfoList
	resp.LabInfos = labInfoList
	resp.ContestInfo = *contest
	err = summaryAndSort(submitGroupData, resp)
	if err != nil {
		appG.RespErr(e.ERROR, fmt.Sprintf("summary dashboard failed [%#v]", err))
		return
	}
	appG.RespSucc(resp)

}

func summaryAndSort(groupData []models.SubmitGroupData, resp *dashboardResp) error {
	if resp.RankDataMap == nil {
		resp.RankDataMap = make(map[uint64]map[uint64]dashboardUserData)
	}

	for _, v := range groupData {
		if _, ok := resp.RankDataMap[v.CreatorId]; !ok {
			resp.RankDataMap[v.CreatorId] = make(map[uint64]dashboardUserData)
		}
		tmp := resp.RankDataMap[v.CreatorId][v.LabID]
		if v.Status == models.LABSUBMITSTATUS_ACCEPTED {
			tmp.IsAc = true
		} else {
			tmp.TimeSum += models.PENAL_TIME
		}
		tmp.SubmitTimes += v.Cnt
		resp.RankDataMap[v.CreatorId][v.LabID] = tmp
	}
	return nil
}

func checkParams(req *dashboardReq, contest *models.Contest, appG *app.Gin) bool {
	if req.ContestId == 0 {
		appG.RespErr(e.INVALID_PARAMS, "please check contest id")
		return false
	}

	// check exists
	contests, err := contest.GetListById(req.ContestId, models.STATUS_ENABLE)
	if err != nil || len(contests) == 0 {
		appG.RespErr(e.INVALID_PARAMS, "contest id not exist")
		return false
	}
	contest = contests[0]
	return true
}