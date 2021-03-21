package contest

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
)

type contestUserReq struct {
	ContestId uint64 `json:"contest_id"`
	Pager models.Pager
}

type contestUserResp struct {
	UserList []*models.ContestUserMap `json:"user_list"`
	Count int `json:"count"`
}

func Users(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	req := &contestUserReq{}
	err := c.BindJSON(req)
	if err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, nil)
		return
	}

	resp := &contestUserResp{}
	contestUserMap := models.ContestUserMap{}
	resp.UserList, err = contestUserMap.GetList(req.Pager, models.STATUS_ALL)
	resp.Count, err = models.GetCountByStatus(models.TABLE_CONTEST_USER_MAP, models.STATUS_ALL)
	if err != nil {
		log.Printf("Get contest from db error [%v]", err)
		appG.RespErr(e.ERROR, "get contest from db error")
		return
	}

	appG.RespSucc(resp)
}


