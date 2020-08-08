package judger

import (
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJudger/caroline"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Judge(c *gin.Context) {

	appG := app.Gin{
		C: c,
	}

	submitIdStr, _ := c.GetPostForm("submit_id")
	submitId, _ := strconv.ParseUint(submitIdStr, 10, 64)

	// 获取judge
	_ = caroline.JudgeSubmit(submitId)
	appG.Response(http.StatusOK, e.SUCCESS, nil)

}
