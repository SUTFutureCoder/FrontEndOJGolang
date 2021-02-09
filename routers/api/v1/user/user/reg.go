package user

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Reg(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	user := models.User{}
	prepare(&user, c)

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

func prepare(user *models.User, c *gin.Context) {
	user.Creator, _ = c.GetPostForm("user_name")
	userPassword, _ := c.GetPostForm("user_password")
	passByte, _ := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	user.UserPassword = string(passByte)
	user.UserType = models.USERTYPE_NORMAL
}
