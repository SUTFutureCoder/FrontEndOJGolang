package user

import (
	"FrontEndOJGolang/pkg/app"
	"github.com/gin-gonic/gin"
)

/**
 * 聚合统计结果
 */
func Summary(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	appG.RespSucc(nil)
}