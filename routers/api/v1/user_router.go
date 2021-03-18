package v1

import (
	"FrontEndOJGolang/routers/api/v1/user/judger"
	"FrontEndOJGolang/routers/api/v1/user/lab"
	"FrontEndOJGolang/routers/api/v1/user/lab_submit"
	"FrontEndOJGolang/routers/api/v1/user/tools/file"
	"FrontEndOJGolang/routers/api/v1/user/user"
	"FrontEndOJGolang/routers/api/v1/user/ws"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
	// 实验区
	labGroup := r.Group("/lab")
	labGroup.POST("/list", lab.LabList)
	labGroup.POST("/info", lab.LabInfo)
	labGroup.POST("/list_with_summary", lab.LabListAndSummary)

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
	userGroup.POST("/year_submit_summary", user.YearSubmitSummary)
	userGroup.POST("/day_submits", lab_submit.DaySubmits)
	userGroup.POST("/change_pwd", user.ChangePwd)

	contestGroup := r.Group("/contest")
	contestGroup.POST("/list_with_summary", contest.ListWithSummary)
	contestGroup.POST("/info", contest.Info)
	contestGroup.POST("/sign", contest.Sign)
	contestGroup.POST("/dashboard", contest.Dashboard)

	contestGroup.POST("/get_labs", contest.GetLabs)

	contestGroup.POST("/get_users", contest.GetUsers)

	// 工具
	tool := r.Group("/tool")
	tool.GET("/getfile", file.GetFile)
	tool.POST("/uploadfile", file.UploadFile)

	// websocket
	r.GET("/ws", ws.Ws)
	r.GET("/ws_judger", ws.WsJudger)
}
