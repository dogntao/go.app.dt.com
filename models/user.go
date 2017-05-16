package models

import (
	"fmt"
)

type User struct {
	Id         int64
	User_name  string
	Pass_word  string
	Company_id int64
	Role       int64
	Status     int64
}

type UserInfo struct {
	User
	UserName string
	PassWord string
}

func (u *UserInfo) LoginCheck() {
	var user User
	res := getOne(user, "cms_user", u.UserName)
	fmt.Println(res)
}
