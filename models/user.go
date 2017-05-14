package models

import (
	"fmt"
	"reflect"
	"strings"
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
	// 获取链接
	db := dbStore.GetConn()
	defer dbStore.RetConn(db)
	// 把字段解析到struct

	// 拼装查询字段
	var user User
	t := reflect.TypeOf(user)
	selArr := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		selArr[i] = strings.ToLower(t.Field(i).Name)
	}
	selField := strings.Join(selArr, ",")
	fmt.Println(selField)

	// 接收结果
	user2 := &User{}
	ele := reflect.ValueOf(user2).Elem()
	leng := ele.NumField()
	oneRow := make([]interface{}, leng)
	for i := 0; i < leng; i++ {
		oneRow[i] = ele.Field(i).Addr().Interface()
	}

	sql := fmt.Sprintf("SELECT %s from cms_user where user_name=?", selField)
	stmt, err := db.Prepare(sql)
	checkErr(err)
	row := stmt.QueryRow(u.UserName)

	err = row.Scan(oneRow...)
	checkErr(err)
	fmt.Println(user2)

	return
}
