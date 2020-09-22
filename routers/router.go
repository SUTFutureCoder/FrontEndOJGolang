package routers

import (
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/routers/api/v1/judger"
	"FrontEndOJGolang/routers/api/v1/lab"
	"FrontEndOJGolang/routers/api/v1/lab_submit"
	"FrontEndOJGolang/routers/api/v1/lab_testcase"
	"FrontEndOJGolang/routers/api/v1/tools"
	"FrontEndOJGolang/routers/api/v1/user"
	"github.com/gin-gonic/gin"

	"FrontEndOJGolang/routers/api/v1/testfield"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(app.CORSMiddleware())

	// 实验区
	labGroup := r.Group("/lab")
	labGroup.POST("/add", lab.AddLab)
	labGroup.POST("/list", lab.LabList)
	labGroup.POST("/info", lab.LabInfo)

	labTestcaseGroup := r.Group("/lab_testcase")
	labTestcaseGroup.POST("/add", lab_testcase.Add)

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

	// 测试区
	test := r.Group("/test")
	test.POST("/screenshot", testfield.ScreenShot)

	// 工具
	tool := r.Group("/tool")
	tool.GET("/getfile", tools.GetFile)
	tool.POST("/uploadfile", tools.UploadFile)

	return r
}
