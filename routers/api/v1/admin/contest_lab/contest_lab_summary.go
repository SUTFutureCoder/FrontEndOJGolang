package contest_lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

type getLabsReq struct {
	ContestId uint64 `json:"contest_id"`
}
type getLabsResp struct {
	ContestLabs []models.Lab `json:"contest_labs"`
}

func GetLabs(c *gin.Context) {
	appG := app.Gin{C: c}

	req := &getLabsReq{}
	err := c.BindJSON(req)
	if err != nil {
		appG.RespErr(e.PARSE_PARAM_ERROR, nil)
		return
	}

	if req.ContestId == 0 {
		appG.RespErr(e.INVALID_PARAMS, "contest id equals 0")
		return
	}

	resp := &getLabsResp{
		ContestLabs: []models.Lab{},
	}
	contestLab := &models.ContestLabMap{}
	_, labIdList, err := contestLab.GetIdMap([]interface{}{req.ContestId}, models.STATUS_ALL)
	if len(labIdList) == 0 {
		appG.RespSucc(resp)
		return
	}

	lab := &models.Lab{}
	labInfo := lab.GetByIds(labIdList)

	// sort
	labInfo = lab.SortLabs(models.ConvertInterfaceToUint64(labIdList), labInfo)

	for _, l := range labInfo {
		resp.ContestLabs = append(resp.ContestLabs, l)
	}
	appG.RespSucc(resp)
}