package app

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/setting"
	"encoding/gob"
	"errors"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var store = sessions.NewCookieStore([]byte(setting.SessionSetting.Token))

type UserSession struct {
	Id       uint64
	Name     string
	UserType int8
}

const SESSIONKEY = "data"

func GetSession(c *http.Request) (*sessions.Session, error) {
	// record to session
	gob.Register(UserSession{})
	session, err := store.Get(c, setting.SessionSetting.SessionUser)
	if err != nil {
		log.Printf("[ERROR] get session store error err[%v] ", err)
		return nil, err
	}
	return session, nil
}

func SetSession(c *http.Request, w http.ResponseWriter, user *models.User) error {
	gob.Register(UserSession{})
	session, err := GetSession(c)
	if err != nil {
		return err
	}
	userSession := UserSession{
		Id:       user.ID,
		Name:     user.Creator,
		UserType: user.UserType,
	}
	session.Values[SESSIONKEY] = userSession
	err = session.Save(c, w)
	if err != nil {
		log.Printf("[ERROR] save session store error user[%v] err[%v] ", user, err)
		return err
	}
	return nil
}

func ExpireSession(c *http.Request, w http.ResponseWriter) error {
	session, err := GetSession(c)
	if err != nil {
		log.Printf("[ERROR] expire session when get session error err[%v] ", err)
		return err
	}
	session.Options.MaxAge = -1
	err = session.Save(c, w)
	if err != nil {
		log.Printf("[ERROR] expire session when exec session error err[%v] ", err)
		return err
	}
	return nil
}

func GetUserFromSession(c *http.Request) (UserSession, error) {
	session, err := GetSession(c)
	if err != nil {
		return UserSession{}, err
	}

	userSession, parseOk := session.Values[SESSIONKEY].(UserSession)
	if !parseOk {
		//log.Printf("[ERROR] parse user info error err[%v]", err)
		return UserSession{}, errors.New("parse user info error")
	}

	if userSession.Id == 0 {
		return UserSession{}, errors.New("unlogin")
	}
	return userSession, nil
}
