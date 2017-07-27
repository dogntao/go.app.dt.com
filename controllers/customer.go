package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"go.app.dt.com/models"
)

var customerModel = &models.Customer{}

type CustomerController struct {
	BaseController
}

func (c *CustomerController) Add() {
	if req.Method == "GET" {
		// 获取id参数编辑
		id, ok := paramMap["id"]
		customerInfoMap := make(map[string]string, 0)
		if ok {
			customerInfoMap = customerModel.Info(id)
		}
		customerInfo, _ := json.Marshal(customerInfoMap)
		assign["Info"] = string(customerInfo)
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
		// 搜索数据
		search := paramMap["search"]
		// 处理分页
		var page int64
		page = 1
		getPage, ok := paramMap["page"]
		if ok {
			page, _ = strconv.ParseInt(getPage, 10, 64)
		}

		total, list := customerModel.Manage(search, page, 10)
		listByte, _ := json.Marshal(list)
		// 传递参数
		assign["Total"] = math.Ceil(float64(total) / 10)
		assign["List"] = string(listByte)
		assign["Search"] = search
		assign["Page"] = page
		c.DisplayAdmin("views/customer/manage.html")
	} else {
	}
}
