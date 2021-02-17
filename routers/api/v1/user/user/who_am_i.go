package user

import (
	"FrontEndOJGolang/pkg/app"
	"github.com/gin-gonic/gin"
)

type whoamiResp struct {
	Id	uint64 `json:"id"`
	Name string `json:"name"`
	UserType int8 `json:"user_type"`
}
func WhoAmI(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	user := app.GetUserFromSessionNoRespErr(appG)
	appG.RespSucc(whoamiResp{
		Id: user.Id,
		Name: user.Name,
		UserType: user.UserType,
	})
}
