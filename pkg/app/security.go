package app

import (
	"FrontEndOJGolang/models"
	"time"
)

/**
 * 限制重复提交频率
 */
func LimitUserSubmitFluency(userId uint64) bool {
	labSubmit, err := models.GetUserLastSubmit(userId)
	if err != nil || labSubmit.ID == 0 {
		return false
	}

	if time.Now().UnixNano()/1e6-labSubmit.CreateTime < 1000*60 {
		return true
	}
	return false
}
