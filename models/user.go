package models

import (
	"fmt"
	"reflect"
)

type User struct {
	Id        int64
	UserName  string
	PassWord  string
	CompanyId int64
	Role      int64
	Status    int64
}

type UserInfo struct {
	User
	UserName string
	PassWord string
}

func (u *UserInfo) LoginCheck() {
	db := dbStore.GetConn()
	defer dbStore.RetConn(db)

	stmt, err := db.Prepare("SELECT * from cms_user where user_name=?")
	checkErr(err)
	row := stmt.QueryRow(u.UserName)
	// 把字段解析到struct
	user := &User{}
	ele := reflect.ValueOf(user).Elem()
	leng := ele.NumField()
	oneRow := make([]interface{}, leng)
	for i := 0; i < leng; i++ {
		oneRow[i] = ele.Field(i).Addr().Interface()
	}

	err = row.Scan(oneRow...)
	checkErr(err)
	// result = append(result, ele.Interface())
	fmt.Println(user.CompanyId)

	return
}
