package routers

import (
	"github.com/gin-gonic/gin"

	"FrontEndOJGolang/routers/api/testfield"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 测试区
	test := r.Group("/test")
	test.POST("/screenshot", testfield.ScreenShot)

	return r
}


