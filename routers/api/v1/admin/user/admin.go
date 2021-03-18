package user

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type adminUserReq struct {
	UserId uint64 `json:"user_id"`
}

func DisableUser(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	var req adminUserReq
	err := c.BindJSON(&req)
	if err != nil {
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	user := &models.User{
		Model: models.Model{ID: req.UserId},
	}
	if !user.ModifyStatus(models.STATUS_DISABLE) {
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}
	appG.RespSucc(nil)
}

func EnableUser(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	var req adminUserReq
	err := c.BindJSON(&req)
	if err != nil {
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	user := &models.User{
		Model: models.Model{ID: req.UserId},
	}
	if !user.ModifyStatus(models.STATUS_ENABLE) {
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}
	appG.RespSucc(nil)
}

type createUserReq struct {
	UserName   string `json:"user_name"`
	UserPasswd string `json:"user_password"`
	UserType   int8   `json:"user_type"`
}

func CreateUser(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		return
	}

	var req createUserReq
	var user models.User
	err := c.BindJSON(&req)
	if err != nil {
		appG.RespErr(e.INVALID_PARAMS, err.Error())
		return
	}
	passByte, _ := bcrypt.GenerateFromPassword([]byte(req.UserPasswd), bcrypt.DefaultCost)
	user.Creator = req.UserName
	user.UserPassword = string(passByte)
	user.UserType = req.UserType
	user.CreatorId = userSession.Id

	// check if user have exist
	exist, err := user.CheckExist()
	if err != nil || exist {
		appG.RespErr(e.INVALID_PARAMS, "username exist")
		return
	}

	err = user.Insert()
	if err != nil {
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}
	appG.RespSucc(nil)
}

type changePasswdReq struct {
	UserId    uint64 `json:"user_id"`
	NewPasswd string `json:"new_user_passwd"`
}

func ChangePasswd(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	userSession := app.GetUserFromSession(appG)
	if userSession.Id != 0 {
		return
	}

	var req changePasswdReq
	c.BindJSON(&req)

	var user models.User
	user.ID = req.UserId
	passByte, _ := bcrypt.GenerateFromPassword([]byte(req.NewPasswd), bcrypt.DefaultCost)
	user.UserPassword = string(passByte)

	_, err := user.UpdatePwd()
	if err != nil {
		log.Printf("update user passwd error[%#v]", err)
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	appG.RespSucc(nil)
}

type grantPermissionReq struct {
	UserId   uint64 `json:"user_id"`
	UserType int8   `json:"user_type"`
}

func GrantPermission(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	usersession := app.GetUserFromSession(appG)
	if usersession.Id != 0 {
		return
	}

	var req grantPermissionReq
	var user models.User
	user.UserType = req.UserType
	user.ID = req.UserId
	if !user.GrantType() {
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}
	appG.RespSucc(nil)
}

type modifyUserReq struct {
	UserId       uint64 `json:"user_id"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	UserType     int8   `json:"user_type"`
}

func ModifyUser(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}
	var req modifyUserReq
	err := c.BindJSON(&req)
	if err != nil {
		appG.RespErr(e.INVALID_PARAMS, nil)
		return
	}

	userSession := app.GetUserFromSession(appG)
	if userSession.Id == 0 {
		return
	}

	user := &models.User{}
	// check user exists
	user.GetById(req.UserId)
	if user.ID == 0 {
		appG.RespErr(e.INVALID_PARAMS, "user not exists")
		return
	}

	// overwrite
	if req.UserType != user.UserType {
		user.UserType = req.UserType
	}
	if req.UserName != user.Creator {
		user.Creator = req.UserName
	}
	if req.UserPassword != "" {
		// get new password hash
		passByte, _ := bcrypt.GenerateFromPassword([]byte(req.UserPassword), bcrypt.DefaultCost)
		req.UserPassword = string(passByte)
		if req.UserPassword != user.UserPassword {
			user.UserPassword = req.UserPassword
		}
	}
	if !user.Modify() {
		appG.RespErr(e.INVALID_PARAMS, "failed to modify user")
		return
	}
	appG.RespSucc(nil)
}
