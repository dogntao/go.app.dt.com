package models

import (
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	id        int64
	user_name string
	pass_word string
	role      int64
	status    int64
}

type UserInfo struct {
	User
	UserName string
	PassWord string
}

var companyTable = "cms_company"

func (u *UserInfo) LoginCheck() (check bool) {
	var user User
	con := "user_name=?"
	bind := []string{u.UserName}
	err := Dtsql.Query(user, companyTable, con, bind)
	checkErr(err)
	err = Dtsql.FetchRow()
	checkErr(err)
	// 验证用户名与密码
	h := md5.New()
	h.Write([]byte(u.PassWord))
	b := h.Sum(nil)
	passMd5 := hex.EncodeToString(b)

	check = false
	if Dtsql.RetMap["pass_word"] == passMd5 {
		check = true
	}
	return
}
