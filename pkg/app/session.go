package app

// 废弃
import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/setting"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

const USERSESSION = "USERSESSION"

type UserSession struct {
	Id       uint64
	Name     string
	UserType int8
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

func GetUserFromSession(appG Gin) UserSession {
	session := sessions.Default(appG.C).Get(setting.SessionSetting.SessionUser)
	if session == nil {
		log.Printf("get session nil")
		appG.RespErr(e.NOT_LOGINED, nil)
		appG.C.Abort()
		return UserSession{}
	}

	userSession, parseOk := session.(UserSession)
	if !parseOk {
		log.Printf("[ERROR] parse user info error err session[%#v]", session)
		appG.RespErr(e.NOT_LOGINED, nil)
		appG.C.Abort()
		return UserSession{}
	}

	if userSession.Id == 0 {
		appG.RespErr(e.NOT_LOGINED, nil)
		appG.C.Abort()
		return UserSession{}
	}
	appG.C.Set(USERSESSION, userSession)
	return userSession
}

func CheckUserAdmin(c *gin.Context) {
	appG := Gin{
		C: c,
	}
	session := GetUserFromSession(appG)
	if session.Id == 0 || session.UserType != models.USERTYPE_ADMIN {
		appG.RespErr(e.UNAUTHORIZED, nil)
		return
	}
	c.Set(USERSESSION, session)
	return
}