package user

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type changePwdReq struct {
	NewPassWd string `json:"new_password"`
}

func ChangePwd(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		return
	}

	var req changePwdReq
	c.BindJSON(&req)

	var user models.User
	user.ID = userSession.Id
	passByte, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassWd), bcrypt.DefaultCost)
	user.UserPassword = string(passByte)

	_, err := user.UpdatePwd()
	if err != nil {
		log.Printf("update user passwd error[%#v]", err)
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	appG.RespSucc(nil)
}
