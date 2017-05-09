package models

import (
	"fmt"
)

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

	stmt, err := db.Prepare("SELECT * from cms_user where user_name=?")
	checkErr(err)
	rows, _ := stmt.Query(u.UserName)
	cols, _ := rows.Columns()
	fmt.Println(cols)

	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))
	dest := make([]interface{}, len(cols))
	for i, _ := range rawResult {
		dest[i] = &rawResult[i]
	}
	if rows.Next() {
		err = rows.Scan(dest...)
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = ""
			} else {
				result[i] = string(raw)
			}
		}
	} else {
	}

	fmt.Println(result)

	return
}
