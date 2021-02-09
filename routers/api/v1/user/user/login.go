package user

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type LoginResp struct {
}

type LoginReq struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

func Login(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	var loginReq LoginReq
	if err := c.BindJSON(&loginReq); err != nil {
		log.Println(err)
	}

	user := new(models.User)
	user.Creator = loginReq.UserName
	err := user.GetByName()
	if err != nil {
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(loginReq.UserPassword))
	if err != nil {
		log.Printf("[ERROR] check password error user[%v] err[%v] ", user, err)
		appG.RespErr(e.INVALID_PARAMS, "please check your password")
		return
	}

	// save session
	err = app.SetSession(c, user)
	if err != nil {
		appG.RespErr(e.ERROR, err)
	}

	appG.RespSucc(user)
}