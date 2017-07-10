package controllers

import (
	"encoding/json"
	"fmt"

	"go.app.dt.com/models"
)

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
		c.DisplayAdmin("views/customer/add.html")
	} else {
		cus := make(map[string]interface{}, 0)
		cus["name"] = req.PostFormValue("name")
		cus["mobile"] = req.PostFormValue("mobile")
		cus["address"] = req.PostFormValue("address")
		cus["balance"] = req.PostFormValue("balance")
		cus["discount"] = req.PostFormValue("discount")
		cus["company_id"] = 1
		cus["is_delete"] = 0
		lastID, err := customerModel.Add(cus)

		jr := &jsonResult{}
		if err != nil {
			jr.Code = 201
			jr.Message = "添加客户失败"
		} else {
			jr.Code = 200
			jr.Message = string(lastID)
		}
		r, _ := json.Marshal(jr)
		fmt.Fprintln(rep, string(r))
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
		c.DisplayAdmin("views/customer/manage.html")
		customerModel.Manage("董涛")
	} else {
	}
}
