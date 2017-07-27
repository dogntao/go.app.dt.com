package models

import (
	"fmt"
	"strconv"
)

type Customer struct {
}

type customerCount struct {
	count string
}

type CustomerInfo struct {
	id         int64
	name       string
	mobile     string
	address    string
	balance    string
	discount   string
	company_id int64
	is_delete  int64
}

var cusTable = "cms_customer"

// 查询详情
func (c *Customer) Info(id string) (info map[string]string) {
	var customerInfo CustomerInfo
	con := "id=?"
	bind := []string{id}
	err := Dtsql.Query(customerInfo, cusTable, con, bind)
	err = Dtsql.FetchRow()
	checkErr(err)
	info = Dtsql.RetMap
	return
}

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

// 列表(返回总数和列表)
func (c *Customer) Manage(seaStr string, pageIndex, pageSize int64) (total int, list []map[string]string) {
	con := ""
	bind := []string{}
	if seaStr != "" {
		con = "(id like ? OR name like ? OR mobile like ? OR address like ?)"
		for i := 0; i < 4; i++ {
			seaStr = "%" + seaStr + "%"
			bind = append(bind, seaStr)
		}
	} else {
		// bind = append(bind, "")
	}
	// 查询总数
	var cusCount customerCount
	err := Dtsql.Query(cusCount, cusTable, con, bind)
	err = Dtsql.FetchRow()
	total, _ = strconv.Atoi(Dtsql.RetMap["count"])

	// 查询列表
	var cusInfo CustomerInfo
	if con == "" {
		con = "1=1"
	}
	con = con + " LIMIT ?,?"
	bind = append(bind, fmt.Sprintf("%d", (pageIndex-1)*pageSize))
	bind = append(bind, fmt.Sprintf("%d", pageSize))
	err = Dtsql.Query(cusInfo, cusTable, con, bind)
	err = Dtsql.FetchAll()
	list = Dtsql.RetRows

	checkErr(err)
	return
}
