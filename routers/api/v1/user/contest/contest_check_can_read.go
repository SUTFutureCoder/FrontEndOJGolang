package contest

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/utils"
	"github.com/gin-gonic/gin"
)

type tryAccessReq struct {
	ContestId uint64 `json:"contest_id"`
}

func TryAccess(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	req := &tryAccessReq{}
	err := c.BindJSON(req)
	if err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, err)
		return
	}

	contest := &models.Contest{}
	contests, err := contest.GetListById(req.ContestId, models.STATUS_ENABLE)
	if len(contests) == 0 {
		appG.RespErr(e.INVALID_PARAMS, "Contest Was Missing")
		return
	}

	userSession := app.GetUserFromSessionNoRespErr(appG)
	if userSession.UserType == models.USERTYPE_ADMIN {
		appG.RespSucc(true)
		return
	}

	reason := checkAccess(contests[0], userSession.Id)
	if reason != "" {
		appG.RespErr(e.CONTEST_ACCESS_DENIED, reason)
		return
	}

	appG.RespSucc(true)
}

func checkAccess(contest *models.Contest, userId uint64) string {
	if contest.ContestStartTime > utils.GetMillTime() {
		return "Contest Not Start"
	}
	if contest.ContestEndTime < utils.GetMillTime() {
		return "Contest Was Ended"
	}
	contestUserMap := &models.ContestUserMap{
		ContestId: contest.ID,
		Model : models.Model{
			CreatorId: userId,
			Status: models.STATUS_ENABLE,
		},
	}
	// check user have signed
	if !contestUserMap.CheckUserSignIn() {
		return "Please Signin The Contest First"
	}

	return ""
}
