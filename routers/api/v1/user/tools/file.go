package tools

import (
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/setting"
	"FrontEndOJGolang/routers/api/v1/user/tools/file"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// be ware of backup upload files
func UploadFile(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		return
	}

	fileReader, err := c.FormFile("file")

	fileTool, err := file.GetFileToolWithUser(setting.ToolSetting.FileToolType, userSession.Id)
	if fileTool == nil || err != nil {
		log.Printf("[ERROR] upload file error:[%v]", err)
		appG.RespErr(e.INVALID_PARAMS, "upload file error")
		return
	}
	filePath, err := fileTool.Put(fileReader)
	if err != nil {
		log.Printf("[ERROR] save upload file error:[%v]", err)
		appG.RespErr(e.INVALID_PARAMS, "save upload file error")
		return
	}
	appG.RespSucc(filePath)
}

func GetFile(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	fileParam := c.Query("file")
	if fileParam == "" || strings.Contains(fileParam, "..") {
		appG.RespErr(e.INVALID_PARAMS, "invalid file url")
		return
	}

	fileTool, err := file.GetFileTool(setting.ToolSetting.FileToolType)
	if fileTool == nil || err != nil {
		log.Printf("[ERROR] upload file error:[%v]", err)
		appG.RespErr(e.INVALID_PARAMS, "upload file error")
		return
	}

	bytes, err := fileTool.Get(setting.ToolSetting.FileBaseDir + "/" + fileParam)
	c.Data(e.SUCCESS, http.DetectContentType(bytes), bytes)
}
