package user

import (
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
)

func WhoAmI(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	user, err := app.GetUserFromSession(c)
	if err != nil || user.Id == 0 {
		appG.RespErr(e.NOT_LOGINED, "please login")
		return
	}

	appG.RespSucc(user)
}
