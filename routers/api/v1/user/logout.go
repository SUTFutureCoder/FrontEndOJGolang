package user

import (
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logout(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	err := app.ExpireSession(c.Request, c.Writer)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
