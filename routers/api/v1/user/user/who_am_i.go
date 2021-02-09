package user

import (
	"FrontEndOJGolang/pkg/app"
	"github.com/gin-gonic/gin"
)

func WhoAmI(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	user := app.GetUserFromSession(appG)
	if user.Id == 0 {return}
	appG.RespSucc(user)
}
