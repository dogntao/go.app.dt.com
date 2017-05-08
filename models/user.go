package models

import "fmt"

type User struct {
	UserName string
	PassWord string
}

func (u *User) LoginCheck() {
	db := dbStore.GetConn()
	defer dbStore.RetConn(db)
	fmt.Println(db)
	fmt.Println("con")
	return
}
