package judger

import (
	"FrontEndOJGolang/pkg/app"
	"github.com/gin-gonic/gin"
)

func Judge(c *gin.Context) {

	appG := app.Gin{
		C: c,
	}

	//submitIdStr, _ := c.GetPostForm("submit_id")
	//submitId, _ := strconv.ParseUint(submitIdStr, 10, 64)

	// 获取judge
	//caroline.JudgeSubmit(submitId)
	appG.RespSucc(nil)

}
