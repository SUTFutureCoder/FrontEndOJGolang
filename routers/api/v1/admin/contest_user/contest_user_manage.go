package contest_user

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type modifyUsersStatusReq struct {
	ContestId uint64 `json:"contest_id"`
	UserIds []interface{} `json:"user_id"`
	Status int `json:"status"`
}

func ModifyUsersStatus(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	req := &modifyUsersStatusReq{}
	if err := c.BindJSON(req); err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, nil)
		return
	}

	if len(req.UserIds) == 0 || req.ContestId == 0 || (req.Status != models.STATUS_ENABLE && req.Status != models.STATUS_DISABLE) {
		appG.RespErr(e.INVALID_PARAMS, "param can not be empty or 0")
		return
	}

	contestUserMap := &models.ContestUserMap{
		ContestId: req.ContestId,
	}
	if !contestUserMap.ModifyStatus(req.UserIds, req.Status) {
		appG.RespErr(e.ERROR, "update contest user map db error")
		return
	}
	appG.RespSucc(nil)
}

type addUsersReq struct {
	ContestId uint64 `json:"contest_id"`
	UserIds []uint64 `json:"user_ids"`
}

func AddUsers(c *gin.Context) {
	appG := app.Gin {
		C: c,
	}

	req := &addUsersReq{}
	if err := c.BindJSON(req); err != nil {
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	if req.ContestId == 0 || len(req.UserIds) == 0 {
		appG.RespErr(e.INVALID_PARAMS, "contest id equals 0 or user ids empty")
		return
	}

	contestUserMap := &models.ContestUserMap {
		ContestId: req.ContestId,
		Model: models.Model{
			Status: models.STATUS_ENABLE,
		},
	}

	for _, v := range req.UserIds {
		// GET USER INFO
		user := &models.User{}
		user.GetById(v)
		if user.ID == 0 {
			appG.RespErr(e.INVALID_PARAMS, "id=" +strconv.FormatUint(v, 10)+ " user not found")
			return
		}
		contestUserMap.CreatorId = user.ID

		// check exists
		if contestUserMap.CheckUserExists() {
			continue
		}

		contestUserMap.Creator = user.Creator
		contestUserMap.CreateTime = utils.GetMillTime()
		insertId, err := contestUserMap.Insert()
		if insertId == 0 || err != nil {
			log.Printf("insert into contest_user_map error [%#v] data[%#v]", err, contestUserMap)
			appG.RespErr(e.ERROR, "insert into db error ")
			return
		}
	}

	appG.RespSucc(nil)

}