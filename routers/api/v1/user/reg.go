package user

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Reg(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	// check if login
	userSession, err := app.GetUserFromSession(c)
	if userSession.Id != 0 {
		appG.Response(http.StatusForbidden, e.INVALID_PARAMS, "You have been login")
		return
	}

	user := models.User{}
	prepare(&user, c)

	// check if user have exist
	exist, err := user.CheckExist()
	if err != nil || exist {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, "username exist")
		return
	}

	err = user.Insert()
	if err != nil {
		log.Printf("[ERROR] insert to user table user[%v] error [%v]", user, err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func prepare(user *models.User, c *gin.Context) {
	user.Creator, _ = c.GetPostForm("user_name")
	userPassword, _ := c.GetPostForm("user_password")
	passByte, _ := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	user.UserPassword = string(passByte)
	user.UserType = models.USERTYPE_NORMAL
}
