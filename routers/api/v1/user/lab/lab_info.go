package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"errors"
	"github.com/gin-gonic/gin"
)

type respLabInfo struct {
	LabInfo models.Lab `json:"lab_info"`
}

type reqLabInfo struct {
	Id uint64 `json:"id" from:"id"`
}

/**
 * ALL IN ONE LABINFO TO PREVERT USER HACK TO SEE LABS
 */
func LabInfo(c *gin.Context) {
	appGin := app.Gin{
		C: c,
	}

	userSession := app.GetUserFromSession(appGin)
	if userSession.Id == 0 {
		return
	}

	var resp respLabInfo
	var req reqLabInfo
	err := c.BindJSON(&req)
	if err != nil {
		appGin.RespErr(e.PARSE_PARAM_ERROR, err)
		return
	}
	resp.LabInfo.ID = req.Id
	lab := &models.Lab{}
	err = lab.GetFullInfo(resp.LabInfo.ID)
	if err != nil {
		appGin.RespErr(e.ERROR, err)
		return
	}
	resp.LabInfo = *lab
	
	/**
	 * check something
	 * 1 Hide lab_sample if user was not admin
	 * 2 Only show status = 1 if user was not admin AND not attend in contest
	 */
	err = checkUserStatus(&userSession, &resp.LabInfo)
	if err != nil {
		appGin.RespErr(e.LAB_INVALID, nil)
		return
	}
	appGin.RespSucc(resp)
}

func checkUserStatus(userSession *app.UserSession, labInfo *models.Lab) error {
	if userSession.UserType == models.USERTYPE_ADMIN {
		return nil
	}

	labInfo.LabSample = ""
	if labInfo.Status != models.STATUS_ENABLE {
		contestLabMap := &models.ContestLabMap{}
		contestUserMap := &models.ContestUserMap{}
		contest := &models.Contest{}

		// get lab attend in contests
		contestIds, err := contestLabMap.GetContestIdsByLabId(labInfo.ID, models.STATUS_ENABLE)
		if err != nil {
			return err
		}

		// filter contest if valid
		contests := contest.GetByIds(contestIds, true)

		// check user if in contest
		var validContestIds []interface{}
		for _, v := range contests{
			validContestIds = append(validContestIds, v.ID)
		}

		if len(validContestIds) == 0 {
			return errors.New("empty valid contest")
		}

		if len(contestUserMap.CheckUserSignInByContestIds(validContestIds, userSession.Id)) == 0 {
			return errors.New("user not signin any contest")
		}

	}
	return nil
}
