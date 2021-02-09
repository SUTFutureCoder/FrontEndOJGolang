package routers

import (
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/setting"
	v1 "FrontEndOJGolang/routers/api/v1"
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.Default()
	gob.Register(app.UserSession{})
	r.Use(sessions.Sessions(setting.SessionSetting.SessionUser, cookie.NewStore([]byte(setting.SessionSetting.Token))))
	r.Use(app.CORSMiddleware())

	v1.InitAdminRouter(r)
	v1.InitUserRouter(r)

	return r
}
