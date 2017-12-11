package model

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	sex      int    `json:"sex"`
	nickname string `json:"nickname"`
}
