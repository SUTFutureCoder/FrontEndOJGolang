package routers

import (
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/setting"
	"FrontEndOJGolang/routers/api/v1/judger"
	"FrontEndOJGolang/routers/api/v1/lab"
	"FrontEndOJGolang/routers/api/v1/lab_submit"
	"FrontEndOJGolang/routers/api/v1/lab_testcase"
	"FrontEndOJGolang/routers/api/v1/testfield"
	"FrontEndOJGolang/routers/api/v1/tools"
	"FrontEndOJGolang/routers/api/v1/user"
	"FrontEndOJGolang/routers/api/v1/ws"
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

	// 实验区
	labGroup := r.Group("/lab")
	labGroup.POST("/list", lab.LabList)
	labGroup.POST("/info", lab.LabInfo)

	labSubmitGroup := r.Group("/lab_submit")
	labSubmitGroup.POST("/submit", lab_submit.Submit)
	labSubmitGroup.POST("/submit_list", lab_submit.SubmitList)
	labSubmitGroup.POST("/submit_list_by_lab_id", lab_submit.SubmitListByLabId)

	judgerGroup := r.Group("/judger")
	judgerGroup.POST("/judge", judger.Judge)
	judgerGroup.POST("/testrun", judger.TestRun)

	userGroup := r.Group("/user")
	userGroup.POST("/reg", user.Reg)
	userGroup.POST("/login", user.Login)
	userGroup.POST("/logout", user.Logout)
	userGroup.POST("/whoami", user.WhoAmI)

	// 测试区
	test := r.Group("/test")
	test.POST("/screenshot", testfield.ScreenShot)

	// 工具
	tool := r.Group("/tool")
	tool.GET("/getfile", tools.GetFile)
	tool.POST("/uploadfile", tools.UploadFile)

	// admin管理区
	admin := r.Group("/admin")
	adminLab := admin.Group("/lab")
	adminLab.POST("/add", lab.AddLab)

	labTestcaseGroup := admin.Group("/lab_testcase")
	labTestcaseGroup.POST("/add", lab_testcase.Add)
	labTestcaseGroup.POST("/list", lab_testcase.List)

	// websocket
	r.GET("/ws", ws.Ws)

	return r
}
