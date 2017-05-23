package models

import "fmt"
import "crypto/md5"
import "encoding/hex"

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
	// 验证用户名与密码
	h := md5.New()
	h.Write([]byte(u.PassWord))
	b := h.Sum(nil)
	passMd5 := hex.EncodeToString(b)

	if mysql.RetMap["pass_word"] == passMd5 {
		fmt.Println("123")
	}
}
