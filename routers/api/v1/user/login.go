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

func Login(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	userName, _ := c.GetPostForm("user_name")
	userPassword, _ := c.GetPostForm("user_password")

	user := new(models.User)
	user.Creator = userName
	err := user.GetByName()
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(userPassword))
	if err != nil {
		log.Printf("[ERROR] check password error user[%v] err[%v] ", user, err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err)
		return
	}

	// save session
	err = app.SetSession(c.Request, c.Writer, user)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err)
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
