package app

// 废弃
import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/setting"
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

type UserSession struct {
	Id       uint64
	Name     string
	UserType int8
}

func UserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func SetSession(c *gin.Context, user *models.User) error {
	session := sessions.Default(c)
	option := sessions.Options{
		MaxAge:   86400 * 7,
		Path:     "/",
		HttpOnly: true,
	}
	userSession := UserSession{
		Id:       user.ID,
		Name:     user.Creator,
		UserType: user.UserType,
	}
	session.Set(setting.SessionSetting.SessionUser, userSession)
	session.Options(option)
	log.Println("test")
	err := session.Save()
	if err != nil {
		log.Printf("[ERROR] save session store error user[%v] err[%v] ", user, err)
		return err
	}
	fmt.Println(session.Get(setting.SessionSetting.SessionUser))
	return nil
}

func ExpireSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Options(sessions.Options{
		MaxAge:   0,
		Path:     "/",
		HttpOnly: true,
	})
	err := session.Save()
	if err != nil {
		log.Printf("[ERROR] expire session when exec session error err[%v] ", err)
		return err
	}
	return nil
}

func GetUserFromSession(c *gin.Context) (UserSession, error) {
	session := sessions.Default(c).Get(setting.SessionSetting.SessionUser)
	if session == nil {
		return UserSession{}, nil
	}

	userSession, parseOk := session.(UserSession)
	if !parseOk {
		//log.Printf("[ERROR] parse user info error err[%v]", err)
		return UserSession{}, errors.New("parse user info error")
	}

	if userSession.Id == 0 {
		return UserSession{}, errors.New("unlogin")
	}
	return userSession, nil
}
