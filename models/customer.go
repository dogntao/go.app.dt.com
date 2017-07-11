package models

import (
	"fmt"
)

type Customer struct {
	count string
}

var cusTable = "cms_customer"

// 新增
func (c *Customer) Add(data map[string]interface{}) (lastId int64, err error) {
	lastId, err = Dtsql.Insert(cusTable, data)
	return
}

// 更新
func (c *Customer) Update(upData map[string]interface{}, conStr string) (affRow int64, err error) {
	affRow, err = Dtsql.Update(cusTable, upData, conStr)
	return
}

// 列表
func (c *Customer) Manage(seaStr string) {
	var customer Customer
	con := ""
	bind := []string{}
	if seaStr != "" {
		con = "(id=? OR name=? OR mobile=? OR address=?)"
		for i := 0; i < 4; i++ {
			bind = append(bind, seaStr)
		}
	} else {
		bind = append(bind, "")
	}
	fmt.Println(bind)
	err := Dtsql.Query(customer, cusTable, con, bind)
	checkErr(err)
	err = Dtsql.FetchAll()
	checkErr(err)
	fmt.Println(Dtsql.RetRows)
}
