package v1

import (
	"FrontEndOJGolang/routers/api/v1/user/judger"
	"FrontEndOJGolang/routers/api/v1/user/lab"
	"FrontEndOJGolang/routers/api/v1/user/lab_submit"
	"FrontEndOJGolang/routers/api/v1/user/tools"
	"FrontEndOJGolang/routers/api/v1/user/user"
	"FrontEndOJGolang/routers/api/v1/user/ws"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
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
	userGroup.POST("/summary", user.Summary)
	userGroup.POST("/change_pwd", user.ChangePwd)

	// 工具
	tool := r.Group("/tool")
	tool.GET("/getfile", tools.GetFile)
	tool.POST("/uploadfile", tools.UploadFile)

	// websocket
	r.GET("/ws", ws.Ws)
}