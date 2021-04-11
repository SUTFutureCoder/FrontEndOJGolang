package lab_submit

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/setting"
	"FrontEndOJGolang/pkg/utils"
	"FrontEndOJGolang/routers/api/v1/user/tools/file"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type submitReq struct {
	LabId      uint64 `json:"lab_id"`
	SubmitData string `json:"submit_data"`
}

func Submit(c *gin.Context) {

	appG := app.Gin{
		C: c,
	}

	var req submitReq
	c.BindJSON(&req)

	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		return
	}

	if app.LimitUserSubmitFluency(userSession.Id) {
		appG.RespErr(e.TOO_MANY_REQUESTS, nil)
		return
	}

	labSubmit := &models.LabSubmit{
		LabID: req.LabId,
		SubmitData: req.SubmitData,
		Model: models.Model{
			CreatorId: userSession.Id,
			Creator: userSession.Name,
			CreateTime: utils.GetMillTime(),
		},
	}
	lastId, err := labSubmit.Insert()
	if err != nil {
		log.Printf("[ERROR] lab submit error[%v]\n", err)
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	appG.RespSucc(lastId)

}

func SubmitWithFile(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	fileReader, err := c.FormFile("file")
	if err != nil {
		appG.RespErr(e.INVALID_PARAMS, err)
		return
	}

	labId, err := strconv.ParseUint(c.GetHeader("data"), 10, 64)
	if labId == 0 || err != nil{
		appG.RespErr(e.INVALID_PARAMS, "header must have lab id")
		return
	}
	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		return
	}

	if app.LimitUserSubmitFluency(userSession.Id) {
		appG.RespErr(e.TOO_MANY_REQUESTS, nil)
		return
	}

	// get url
	fileTool, err := file.GetFileToolWithUser(setting.ToolSetting.FileToolType, userSession.Id)
	if fileTool == nil || err != nil {
		log.Printf("[ERROR] upload file error:[%v]", err)
		appG.RespErr(e.INVALID_PARAMS, "upload file error")
		return
	}
	filePath, err := fileTool.Put(fileReader)

	// insert to submit table
	labSubmit := &models.LabSubmit{
		LabID: labId,
		SubmitData: filePath,
		SubmitType: models.SUBMIT_TYPE_PACKAGE,
		Model: models.Model{
			CreatorId: userSession.Id,
			Creator: userSession.Name,
			CreateTime: utils.GetMillTime(),
		},
	}
	lastId, err := labSubmit.Insert()
	if err != nil {
		log.Printf("[ERROR] lab submit with file error[%v]\n", err)
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}
	appG.RespSucc(lastId)
}
