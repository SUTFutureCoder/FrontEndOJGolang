package user

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func Reg(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	userName, _ := c.GetPostForm("user_name")
	userPassword, _ := c.GetPostForm("user_password")
	passByte, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[ERROR] password gen error: %v", err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	user := models.User{
		Model: models.Model {
			Creator: userName,
			CreateTime: time.Now().UnixNano() / 1e6,
		},
		UserPassword: string(passByte),
		UserType:     models.USERTYPE_NORMAL,
	}
	err = user.Insert()
	if err != nil {
		log.Printf("[ERROR] insert to user table user[%v] error [%v]", user, err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
