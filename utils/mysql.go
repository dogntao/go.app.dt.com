/*
 * @Author: dongtao
 * @Date: 2017-04-21 16:55:04
 * @Desc: mysql 方法封装
 * @Last Modified by: dongtao
 * @Last Modified time: 2017-04-21 17:55:33
 */
package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Conf map[string]string

type Mysql struct {
	host    string
	user    string
	pass    string
	dbname  string
	charset string

	db *sql.DB
}

func Connect(c Conf) (mysql *Mysql, err error) {
	mysql = new(Mysql)
	mysql.host = c["host"]
	mysql.user = c["user"]
	mysql.pass = c["pass"]
	mysql.dbname = c["dbname"]
	mysql.charset = c["charset"]

	// db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/dt?charset=utf8")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", mysql.user, mysql.pass, mysql.host, mysql.dbname, mysql.charset)
	mysql.db, err = sql.Open("mysql", dsn)
	checkErr(err)

	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
