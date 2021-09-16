package contest

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type signReq struct {
	ContestId uint64 `json:"contest_id"`
}

func Sign(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	req := &signReq{}
	err := c.BindJSON(req)
	if err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, nil)
		return
	}

	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		return
	}

	contestUserMap := &models.ContestUserMap{
		ContestId: req.ContestId,
		Model: models.Model{
			CreatorId: userSession.Id,
			Creator: userSession.Name,
		},
	}
	if contestUserMap.CheckUserSignIn() {
		appG.RespErr(e.INVALID_PARAMS, "User Had Signed This Contest")
		return
	}

	contestUserMap.Status = models.STATUS_ENABLE
	contestUserMap.CreateTime = utils.GetMillTime()
	lastId, err := contestUserMap.Insert()
	if lastId == 0 || err != nil {
		log.Printf("Contest Signin Failed error[%v]", err)
		appG.RespErr(e.ERROR, "Signin Failed")
		return
	}
	appG.RespSucc(lastId)
}


type getSignStatusReq struct {
	ContestIds []uint64 `json:"contest_ids"`
}
type getSignStatusResp struct {
	SignInContestIds []uint64 `json:"sign_in_contest_ids"`
}
func GetContestsSignStatus(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	req := &getSignStatusReq{}
	err := c.BindJSON(req)
	if err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, nil)
		return
	}

	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		return
	}

	contestUserMap := &models.ContestUserMap{}
	signInContestIds := contestUserMap.GetUserSignInByContestIds(models.ConvertUint64ToInterface(req.ContestIds), userSession.Id);
	resp := &getSignStatusResp{
		SignInContestIds: models.ConvertInterfaceToUint64(signInContestIds),
	}
	appG.RespSucc(resp)
}