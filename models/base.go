package models

import "go.app.dt.com/utils"

var mysql = &utils.Mysql{}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
