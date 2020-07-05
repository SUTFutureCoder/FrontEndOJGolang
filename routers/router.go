package routers

import (
	"FrontEndOJGolang/routers/api/v1/judger"
	"FrontEndOJGolang/routers/api/v1/lab"
	"FrontEndOJGolang/routers/api/v1/lab_submit"
	"FrontEndOJGolang/routers/api/v1/lab_testcase"
	"FrontEndOJGolang/routers/api/v1/user"
	"github.com/gin-gonic/gin"

	"FrontEndOJGolang/routers/api/v1/testfield"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 实验区
	labGroup := r.Group("/lab")
	labGroup.POST("/add", lab.AddLab)

	labTestcaseGroup := r.Group("/lab_testcase")
	labTestcaseGroup.POST("/add", lab_testcase.Add)

	labSubmitGroup := r.Group("/lab_submit")
	labSubmitGroup.POST("/submit", lab_submit.Submit)

	judgerGroup := r.Group("/judger")
	judgerGroup.POST("/judge", judger.Judge)

	userGroup := r.Group("/user")
	userGroup.POST("/reg", user.Reg)
	userGroup.POST("/login", user.Login)

	// 测试区
	test := r.Group("/test")
	test.POST("/screenshot", testfield.ScreenShot)

	return r
}
