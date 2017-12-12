package model

import (
	"mkgo/mkdb"
	"mkgo/mklog"
	"go.uber.org/zap"
)

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Sex      int    `json:"sex"`
	Nickname string `json:"nickname"`
}

func (user *User) Add() bool {
	inertUser := `INSERT INTO user ( name, password) VALUES (?, ?)`
	_, err := mkdb.DB.Exec(inertUser, user.Name, user.Password)
	if err != nil {
		mklog.Logger.Debug("[user]", zap.Error(err))
		return false
	}
	return true
}

func GetUserById(id string) *User {
	var user User
	err := mkdb.DB.Get(&user, `SELECT name, password from user where id=? limit 1`, id)
	if err != nil {
		mklog.Logger.Debug("[user]", zap.Error(err))
		return nil
	}
	return &user
}

func GetUserList() []User {
	var users []User
	err := mkdb.DB.Select(&users, "SELECT name, password FROM user")
	if err != nil {
		mklog.Logger.Debug("[user]", zap.Error(err))
		return nil
	}
	return users
}
