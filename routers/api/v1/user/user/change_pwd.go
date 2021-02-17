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
	OldPassWd string `json:"old_password"`
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

	user := new(models.User)
	user.Creator = userSession.Name
	err := user.GetByName()
	if err != nil {
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	// check old password
	err = bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(req.OldPassWd))
	if err != nil {
		log.Printf("[ERROR] check password error user[%v] err[%v] ", user, err)
		appG.RespErr(e.INVALID_PARAMS, "please check your old password")
		return
	}

	user.ID = userSession.Id
	passByte, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassWd), bcrypt.DefaultCost)
	user.UserPassword = string(passByte)

	_, err = user.UpdatePwd()
	if err != nil {
		log.Printf("update user passwd error[%#v]", err)
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	appG.RespSucc(nil)
}
