// Package main Auto generated.
package models

import "log"

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
