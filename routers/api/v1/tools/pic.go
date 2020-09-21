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
func UploadPic(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	picFileReader, err := c.FormFile("img")
	if err != nil {
		log.Printf("[ERROR] upload file error:[%v]", err)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, "upload file error")
		return
	}
	err = c.SaveUploadedFile(picFileReader, setting.ToolSetting.PicBaseDir+"/"+picFileReader.Filename)
	if err != nil {
		log.Printf("[ERROR] save upload file error:[%v]", err)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, "save upload file error")
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, picFileReader.Filename)
}

func GetPic(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	pic := c.Query("pic")
	if pic == "" || strings.Contains(pic, "..") || strings.Contains(pic, "/") {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, "invalid pic url")
		return
	}
	c.File(setting.ToolSetting.PicBaseDir + "/" + pic)
}
