package models

import "reflect"
import "strings"
import "fmt"
import "database/sql"

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
	// 拼装查询字段
	t := reflect.TypeOf(user)
	filedArr := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		filedArr[i] = t.Field(i).Name
	}
	field := strings.Join(filedArr, ",")

	// 获取链接
	db := dbStore.GetConn()
	defer dbStore.RetConn(db)
	selSql := fmt.Sprintf("SELECT %s FROM cms_user WHERE user_name=?", field)
	stmt, err := db.Prepare(selSql)
	checkErr(err)
	rows, err := stmt.Query(u.UserName)
	checkErr(err)
	if rows.Next() {
		cols, err := rows.Columns()
		checkErr(err)
		scanArgs := make([]interface{}, len(cols))
		values := make([]sql.RawBytes, len(cols))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		err = rows.Scan(scanArgs...)
		checkErr(err)
		retMap := make(map[string]string)
		for i := 0; i < len(cols); i++ {
			retMap[cols[i]] = string(values[i])
		}
		fmt.Println(retMap)
	}

}
