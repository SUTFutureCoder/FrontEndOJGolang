package tools

import (
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/setting"
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
	fileReader, err := c.FormFile("file")
	if err != nil {
		log.Printf("[ERROR] upload file error:[%v]", err)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, "upload file error")
		return
	}
	err = c.SaveUploadedFile(fileReader, setting.ToolSetting.FileBaseDir+"/"+fileReader.Filename)
	if err != nil {
		log.Printf("[ERROR] save upload file error:[%v]", err)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, "save upload file error")
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, fileReader.Filename)
}

func GetFile(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	file := c.Query("file")
	if file == "" || strings.Contains(file, "..") || strings.Contains(file, "/") {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, "invalid file url")
		return
	}
	c.File(setting.ToolSetting.FileBaseDir + "/" + file)
}
