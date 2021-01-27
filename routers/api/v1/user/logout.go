package user

import (
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	err := app.ExpireSession(c)
	if err != nil {
		appG.RespErr(e.ERROR, nil)
	}
	appG.RespSucc(nil)
}
