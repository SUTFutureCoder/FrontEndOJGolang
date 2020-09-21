package judger

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func TestRun(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	code, _ := c.GetPostForm("code")

	if code == "" {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, "empty field code")
		return
	}

	userSession, err := app.GetUserFromSession(c.Request)
	if err != nil {
		log.Printf("[ERROR] get user session error[%v]\n", err)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, "get user session failed")
		return
	}

	// Only > normal user can run test code
	if userSession.UserType <= models.USERTYPE_NORMAL {
		appG.Response(http.StatusOK, e.UNAUTHORIZED, "user can not run test code")
		return
	}

	// Test code only can be run manually, can not be run by auto judger

	// STEP1 build and write fake submit

	// STEP2 build fake testcase id=0

	// STEP3 run and return result

}
