package models

import "go.app.dt.com/utils"

var dbStore = &utils.DbStore{}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getOne() {

}
