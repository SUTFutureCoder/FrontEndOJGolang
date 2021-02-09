package tools

import (
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/setting"
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

// be ware of backup upload files
func UploadFile(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {return}

	fileReader, err := c.FormFile("file")
	if err != nil {
		log.Printf("[ERROR] upload file error:[%v]", err)
		appG.RespErr(e.INVALID_PARAMS, "upload file error")
		return
	}

	// gen new hashed filename
	h := sha1.New()
	h.Write([]byte(fileReader.Filename))
	fileName := hex.EncodeToString(h.Sum(nil)) + path.Ext(fileReader.Filename)

	filePath := setting.ToolSetting.FileBaseDir + "/" + strconv.FormatUint(userSession.Id, 10)
	os.MkdirAll(filePath, os.ModePerm)
	filePath += "/" + fileName

	err = c.SaveUploadedFile(fileReader, filePath)
	if err != nil {
		log.Printf("[ERROR] save upload file error:[%v]", err)
		appG.RespErr(e.INVALID_PARAMS, "save upload file error")
		return
	}
	appG.RespSucc(fileName)
}

func GetFile(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {return}

	file := c.Query("file")
	if file == "" || strings.Contains(file, "..") || strings.Contains(file, "/") {
		appG.RespErr(e.INVALID_PARAMS, "invalid file url")
		return
	}

	filePath := setting.ToolSetting.FileBaseDir + "/" + strconv.FormatUint(userSession.Id, 10) + "/" + file
	c.File(filePath)
}
