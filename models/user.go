package models

import "fmt"
import "reflect"

type User struct {
	id         int64
	UserName   string
	PassWord   string
	pass_word  string
	company_id int64
	role       int64
	status     int64
}

func (u *User) LoginCheck() {
	db := dbStore.GetConn()
	defer dbStore.RetConn(db)

	var user User
	t := reflect.TypeOf(user)
	for k := 0; k < t.NumField(); k++ {
		fmt.Println(t.Field(k).Name)
	}

	stmt, err := db.Prepare("SELECT id,pass_word,company_id,role,status from cms_user where user_name=?")
	checkErr(err)
	rows := stmt.QueryRow(u.UserName)
	err = rows.Scan(&u.id, &u.pass_word, &u.company_id, &u.role, &u.status)
	checkErr(err)
	fmt.Println(u)
	return
}
