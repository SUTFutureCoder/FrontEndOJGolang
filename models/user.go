// Package main Auto generated.
package models

import (
	"FrontEndOJGolang/pkg/utils"
	"log"
)

// User ...
type User struct {
	Model
	// UserPassword 用户密码
	UserPassword string `json:"user_password"`
	// UserType 用户类型
	UserType int8 `json:"user_type"`
}

const (
	USERTYPE_NORMAL = iota
	USERTYPE_ADMIN
	USERTYPE_TEST
)

func (u *User) GetByName() error {
	stmt, err := DB.Prepare("SELECT id, user_password, user_type, status, creator, create_time, update_time FROM user WHERE creator = ?")
	if err != nil {
		log.Printf("[ERROR] prepare sql error user[%v] err[%v]", u, err)
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRow(
		&u.Creator,
	)
	err = row.Scan(
		&u.ID,
		&u.UserPassword,
		&u.UserType,
		&u.Status,
		&u.Creator,
		&u.CreateTime,
		&u.UpdateTime,
	)
	return err
}

func GetUserById(userId uint64) *User {
	var u User
	stmt, err := DB.Prepare("SELECT id, user_password, user_type, status, creator, create_time, update_time FROM user WHERE id=?")
	if err != nil {
		log.Printf("[ERROR] prepare sql error user[%v] err[%v]", u, err)
		return &u
	}
	defer stmt.Close()
	row := stmt.QueryRow(
		&userId,
	)
	err = row.Scan(
		&u.ID,
		&u.UserPassword,
		&u.UserType,
		&u.Status,
		&u.Creator,
		&u.CreateTime,
		&u.UpdateTime,
	)
	return &u
}

func (u *User) Insert() error {
	stmt, err := DB.Prepare("INSERT INTO user(user_password, user_type, creator, create_time) VALUES (?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		log.Printf("[ERROR] database exec error input[%v] err[%v]", u, err)
		return err
	}

	_, err = stmt.Exec(
		u.UserPassword,
		u.UserType,
		u.Creator,
		u.CreateTime,
	)
	return nil
}

func (u *User) CheckExist() (bool, error) {
	stmt, err := DB.Prepare("SELECT COUNT(1) FROM user WHERE creator=?")
	defer stmt.Close()
	if err != nil {
		log.Printf("[ERROR] database exec error input[%v] err[%v]", u, err)
		return false, err
	}
	row := stmt.QueryRow(
		u.Creator,
	)
	i := 0
	err = row.Scan(&i)
	return i > 0, err
}

func (u *User) UpdatePwd() (bool, error) {
	stmt, err := DB.Prepare("UPDATE user SET user_password=?, update_time=? WHERE id=?")
	if err != nil {
		log.Printf("update lab status error [%#v]", err)
		return false, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(u.UserPassword, utils.GetMillTime(), u.ID)
	row, err := res.RowsAffected()
	return row > 0, err
}

type UserSearchParam struct {
	UserId       uint64 `json:"user_id"`
	UserName     string `json:"user_name"`
	UserIdOrName string `json:"user_id_or_name"`
}

func prepareUserSearchParams(searchParam *UserSearchParam, pager *Pager, stmtPrepare *string, querys *[]interface{}) {
	if searchParam.UserId != 0 || searchParam.UserName != "" || searchParam.UserIdOrName != "" {
		*stmtPrepare += " WHERE "
	}
	if searchParam.UserIdOrName != "" {
		*stmtPrepare += " (id=? OR creator=?) "
		*querys = append(*querys, &searchParam.UserIdOrName, &searchParam.UserIdOrName)
		return
	}

	if searchParam.UserId != 0 {
		*stmtPrepare += " id=? "
		*querys = append(*querys, &searchParam.UserId)
	}

	if searchParam.UserName != "" {
		if len(*querys) != 0 {
			*stmtPrepare += " AND "
		}
		*stmtPrepare += " creator=? "
		*querys = append(*querys, &searchParam.UserName)
	}

	DefaultPage(&pager.Page, &pager.PageSize)
	offset := (pager.Page - 1) * pager.PageSize

	*stmtPrepare += " ORDER BY id DESC LIMIT ? OFFSET ?"
	*querys = append(*querys, &pager.PageSize, &offset)

}

func GetUserList(searchParam UserSearchParam, pager Pager) ([]User, error) {
	stmtPrepare := "SELECT id, user_type, status, creator_id, creator, create_time, update_time FROM user "
	var querys []interface{}

	prepareUserSearchParams(&searchParam, &pager, &stmtPrepare, &querys)

	var users []User
	stmt, err := DB.Prepare(stmtPrepare)
	if err != nil {
		log.Printf("get user list status error [%#v]", err)
		return users, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(querys...)
	if err != nil {
		log.Printf("get user list from db error [%v]", err)
		return nil, err
	}

	if rows == nil {
		return nil, err
	}
	var userList []User
	for rows.Next() {
		var user User
		err = rows.Scan(
			&user.ID, &user.UserType, &user.Status, &user.CreatorId, &user.Creator, &user.CreateTime, &user.UpdateTime,
		)
		userList = append(userList, user)
	}
	return userList, err
}

func GetUserCount(searchParam UserSearchParam, pager Pager) (int, error) {
	stmtPrepare := "SELECT count(1) as cnt FROM user "
	var querys []interface{}

	prepareUserSearchParams(&searchParam, &pager, &stmtPrepare, &querys)

	var cnt int
	stmt, err := DB.Prepare(stmtPrepare)
	if err != nil {
		log.Printf("get user count error [%#v]", err)
		return cnt, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(querys...)
	err = row.Scan(&cnt)
	return cnt, err
}

func (u *User) GrantUserType() bool {
	stmt, err := DB.Prepare("UPDATE user SET user_type=?, update_time=? WHERE id=?")
	if err != nil {
		log.Printf("update user type error [%#v]", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.UserType, utils.GetMillTime(), u.ID)
	if err != nil {
		log.Printf("update user type error[%#v]", err)
		return false
	}
	return true
}

func ModifyUserStatus(id uint64, status int) bool {
	stmt, err := DB.Prepare("UPDATE user SET status=?, update_time=? WHERE id=?")
	if err != nil {
		log.Printf("update user status error [%#v]", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(status, utils.GetMillTime(), id)
	if err != nil {
		log.Printf("update user status error[%#v]", err)
		return false
	}
	return true
}

func (u *User) Modify() bool {
	stmt, err := DB.Prepare("UPDATE user SET user_password=?, user_type=?, creator=?, update_time=? WHERE id=?")
	if err != nil {
		log.Printf("update user type error [%#v]", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.UserPassword, u.UserType, u.Creator, utils.GetMillTime(), u.ID)
	if err != nil {
		log.Printf("update user type error[%#v]", err)
		return false
	}
	return true
}
