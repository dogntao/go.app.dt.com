package models

import "fmt"

type User struct {
	id         int64
	user_name  string
	pass_word  string
	company_id int64
	role       int64
	status     int64
}

type UserInfo struct {
	User
	UserName string
	PassWord string
}

func (u *UserInfo) LoginCheck() {
	var user User
	con := "user_name=?"
	err := mysql.Query(user, "cms_user", con, u.UserName)
	checkErr(err)
	err = mysql.FetchOne()
	checkErr(err)
	fmt.Println(mysql.RetMap)
}
