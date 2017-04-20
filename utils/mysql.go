package utils

import (
	"database/sql"
	"fmt"

	"go.app.dt.com/conf"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
}

func Connect(c string) {
	fmt.Println(conf.XMmall)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/dt?charset=utf8")
	checkErr(err)

	// 查询数据
	rows, err := db.Query("select * from cms_company")
	checkErr(err)
	for rows.Next() {
		var id int
		var user_name string
		var pass_word string
		var role int
		var status int
		err = rows.Scan(&id, &user_name, &pass_word, &role, &status)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(user_name)
		fmt.Println(pass_word)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
