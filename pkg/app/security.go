package app

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/utils"
)

/**
 * 限制重复提交频率
 */
func LimitUserSubmitFluency(userId uint64) bool {
	labSubmit := &models.LabSubmit{}
	err := labSubmit.GetUserLastSubmit(userId)
	if err != nil || labSubmit.ID == 0 {
		return false
	}

	if utils.GetMillTime()-labSubmit.CreateTime < 1000*5 {
		return true
	}
	return false
}
