package models

import (
	"fmt"
	"reflect"
	"strings"

	"go.app.dt.com/utils"
)

var dbStore = &utils.DbStore{}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// 注册表和struct
type modStruct struct {
	cms_user User
}

func getOne(user User, tab string, userName string) *User {
	// 获取链接
	db := dbStore.GetConn()
	defer dbStore.RetConn(db)
	// 拼装查询字段
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

	sql := fmt.Sprintf("SELECT %s from %s where id=?", selField, tab)
	stmt, err := db.Prepare(sql)
	checkErr(err)
	row := stmt.QueryRow(userName)

	err = row.Scan(oneRow...)
	checkErr(err)
	// fmt.Println(user2.Company_id)
	return user2

}
