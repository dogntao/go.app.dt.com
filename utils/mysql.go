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
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Conf map[string]string

type Mysql struct {
	host    string
	user    string
	pass    string
	dbname  string
	charset string

	DbStore
	rows   *sql.Rows
	cols   []string
	RetMap map[string]string
}

func Connect(c Conf) (db *sql.DB, err error) {
	mysql := new(Mysql)
	mysql.host = c["host"]
	mysql.user = c["user"]
	mysql.pass = c["pass"]
	mysql.dbname = c["dbname"]
	mysql.charset = c["charset"]

	// db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/dt?charset=utf8")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", mysql.user, mysql.pass, mysql.host, mysql.dbname, mysql.charset)
	db, err = sql.Open("mysql", dsn)
	checkErr(err)
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// query
func (mysql *Mysql) Query(field interface{}, table, con string, bind ...interface{}) (err error) {
	// 拼装查询字段
	t := reflect.TypeOf(field)
	filedArr := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		filedArr[i] = t.Field(i).Name
	}
	fieldString := strings.Join(filedArr, ",")

	// 获取链接
	db := mysql.GetConn()
	defer mysql.RetConn(db)
	selSql := fmt.Sprintf("SELECT %s FROM %s WHERE %s", fieldString, table, con)
	stmt, err := db.Prepare(selSql)
	mysql.rows, err = stmt.Query(bind...)
	return
}

// fetchOne
func (mysql *Mysql) FetchOne() (err error) {
	if mysql.rows.Next() {
		mysql.cols, err = mysql.rows.Columns()
		// fmt.Println(mysql.cols)
		scanArgs := make([]interface{}, len(mysql.cols))
		values := make([]sql.RawBytes, len(mysql.cols))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		err = mysql.rows.Scan(scanArgs...)
		mysql.RetMap = make(map[string]string)
		for i := 0; i < len(mysql.cols); i++ {
			mysql.RetMap[mysql.cols[i]] = string(values[i])
		}
	}
	return
}
