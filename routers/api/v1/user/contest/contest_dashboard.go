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
	UserAcList map[uint64][]uint64 `json:"user_ac_list"`
	UserTimeSumList map[uint64]int `json:"user_time_sum_list"`
}

func Dashboard(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	req := &dashboardReq{}
	resp := &dashboardResp{}
	if err := c.BindJSON(req); err != nil {
		appG.RespErr(e.INVALID_PARAMS, "parse dashboard param error")
		return
	}

	contest := checkParamsAndGetContest(req, &appG)
	if nil == contest {
		return
	}

	contestLabMap := &models.ContestLabMap{}
	_, labIdList, err := contestLabMap.GetIdMap([]interface{}{contest.ID}, models.STATUS_ALL)
	if err != nil {
		log.Printf("get dashboard contest lab map error [%#v] data [%#v]", err, contest)
		appG.RespErr(e.ERROR, "get dashboard contest lab map error")
		return
	}

	labSubmit := &models.LabSubmit{}
	submitGroupData := labSubmit.GroupByUserAndLabIds(contest.ID, labIdList)

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
		resp.UserAcList = make(map[uint64][]uint64)
		resp.UserTimeSumList = make(map[uint64]int)
	}

	for _, v := range groupData {
		if _, ok := resp.RankDataMap[v.CreatorId]; !ok {
			resp.RankDataMap[v.CreatorId] = make(map[uint64]dashboardUserData)
		}
		tmp := resp.RankDataMap[v.CreatorId][v.LabID]
		if v.Status == models.LABSUBMITSTATUS_ACCEPTED {
			resp.UserAcList[v.CreatorId] = append(resp.UserAcList[v.CreatorId], v.LabID)
			tmp.IsAc = true
		} else {
			tmp.TimeSum += models.PENAL_TIME * v.Cnt
			resp.UserTimeSumList[v.CreatorId] += tmp.TimeSum
		}
		tmp.SubmitTimes += v.Cnt
		resp.RankDataMap[v.CreatorId][v.LabID] = tmp
	}
	return nil
}

func checkParamsAndGetContest(req *dashboardReq, appG *app.Gin) *models.Contest {
	if req.ContestId == 0 {
		appG.RespErr(e.INVALID_PARAMS, "please check contest id")
		return nil
	}

	// check exists
	contest := &models.Contest{}
	contests, err := contest.GetListById(req.ContestId, models.STATUS_ENABLE)
	if err != nil || len(contests) == 0 {
		appG.RespErr(e.INVALID_PARAMS, "contest id not exist")
		return nil
	}
	return contests[0]
}