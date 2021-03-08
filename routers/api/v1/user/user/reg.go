package user

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type regReq struct {
	UserName string `json:"user_name"`
	UserPassWord string `json:"user_password"`
}

func Reg(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	var req regReq
	err := c.BindJSON(&req)
	if err != nil || req.UserName == "" || req.UserPassWord == ""{
		appG.RespErr(e.INVALID_PARAMS, "please check your param")
		return
	}

	user := models.User{}
	user.Creator = req.UserName
	passByte, _ := bcrypt.GenerateFromPassword([]byte(req.UserPassWord), bcrypt.DefaultCost)
	user.UserPassword = string(passByte)
	user.UserType = models.USERTYPE_NORMAL

	// check if user have exist
	exist, err := user.CheckExist()
	if err != nil || exist {
		appG.RespErr(e.INVALID_PARAMS, "username exist")
		return
	}

	err = user.Insert()
	if err != nil {
		log.Printf("[ERROR] insert to user table user[%v] error [%v]", user, err)
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}
	appG.RespSucc(nil)
}
