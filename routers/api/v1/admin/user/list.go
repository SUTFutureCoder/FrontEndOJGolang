package user

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
)

type userListReq struct {
	SearchParam models.UserSearchParam `json:"user_search"`
	models.Pager
}

type userListWithSummary struct {
	UserInfo models.User          `json:"user_info"`
	Summary  models.SubmitSummary `json:"summary"`
}

type userListResp struct {
	UserList []userListWithSummary `json:"user_list"`
	Count    int                   `json:"count"`
}

func List(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		return
	}

	var req userListReq
	var resp userListResp
	err := c.BindJSON(&req)
	if err != nil {
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	var users []models.User
	user := &models.User{}
	users, err = user.GetList(req.SearchParam, req.Pager)
	resp.Count, err = user.GetCount(req.SearchParam, req.Pager)

	if err != nil {
		log.Printf("get db list error while get lab list[%#v]", err)
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	var userIds []interface{}
	for _, user := range users {
		userIds = append(userIds, user.ID)
	}

	if len(userIds) == 0 {
		appG.RespSucc(resp)
		return
	}

	labSubmit := &models.LabSubmit{}
	userSubmitSummary := labSubmit.SummaryUserSubmits(userIds)

	// summary
	for _, user := range users {
		var tmpUserListwithsummary userListWithSummary
		tmpUserListwithsummary.UserInfo = user
		if s, ok := userSubmitSummary[user.ID]; ok {
			tmpUserListwithsummary.Summary = *s.UserSubmitSummary
		}
		resp.UserList = append(resp.UserList, tmpUserListwithsummary)
	}
	appG.RespSucc(resp)
}
