package controllers

import "go.app.dt.com/models"
import "fmt"

type addCustomer struct {
	name       string
	mobile     string
	address    string
	balance    string
	discount   string
	company_id int64
	is_delete  int64
}

var customerModel = &models.Customer{}

type CustomerController struct {
	BaseController
}

func (c *CustomerController) Add() {
	if req.Method == "GET" {
		c.Display("views/customer/add.html")
	} else {
		cus := make(map[string]interface{}, 0)
		cus["name"] = "test"
		cus["mobile"] = "18610040137"
		cus["address"] = "河南 禹州"
		cus["balance"] = "600"
		cus["discount"] = "100"
		cus["company_id"] = 1
		cus["is_delete"] = 0

		lastId, err := customerModel.Add(cus)
		fmt.Println(lastId, err)
	}
}

func (c *CustomerController) Update() {
	// if req.Method == "GET" {
	// 	c.Display("views/customer/add.html")
	// } else {
	cus := make(map[string]interface{}, 0)
	cus["name"] = "test56789"
	cus["mobile"] = "18101332343"

	conStr := "id=13"
	affRow, err := customerModel.Update(cus, conStr)
	fmt.Println(affRow, err)
	// }
}

// 列表
func (c *CustomerController) Manage() {
	if req.Method == "GET" {
		c.Display("views/customer/manage.html")
	} else {
	}
}
