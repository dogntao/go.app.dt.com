package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"go.app.dt.com/models"
)

var customerModel = &models.Customer{}

type CustomerController struct {
	BaseController
}

type CustomerSelect2 struct {
	Total int
	Items []map[string]string
}

type CustomerInfoJson struct {
	Code   int
	Result map[string]string
}

// select2
func (c *CustomerController) Select2() {
	if req.Method == "GET" {
		// 搜索数据
		search := paramMap["q"]
		// 处理分页
		var page int64
		page = 1
		getPage, ok := paramMap["page"]
		if ok {
			page, _ = strconv.ParseInt(getPage, 10, 64)
		}

		total, list := customerModel.Manage(search, page, 10)

		customerSelect2 := &CustomerSelect2{}
		customerSelect2.Total = total

		// 返回id和text
		if len(list) > 0 {
			for _, val := range list {
				customerInfo := make(map[string]string, 0)
				customerInfo["id"] = val["id"]
				customerInfo["text"] = val["name"]
				customerSelect2.Items = append(customerSelect2.Items, customerInfo)
			}
		}
		customerSelect2Byte, _ := json.Marshal(customerSelect2)
		fmt.Fprintln(rep, string(customerSelect2Byte))
	}
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
		// assign["Total"] = math.Ceil(float64(total) / 10)
		assign["Total"] = total
		assign["List"] = string(listByte)
		assign["Search"] = search
		assign["Page"] = page
		c.DisplayAdmin("views/customer/manage.html")
	}
}

// 新增/编辑
func (c *CustomerController) Add() {
	// 根据id获取详情
	id, ok := paramMap["id"]
	customerInfoMap := make(map[string]string, 0)
	if ok {
		customerInfoMap, _ = customerModel.Info(id)
	}
	customerInfoByte, _ := json.Marshal(customerInfoMap)
	customerInfo := string(customerInfoByte)

	if req.Method == "GET" {
		// 获取到id参数,编辑
		assign["Id"] = id
		assign["Info"] = customerInfo
		c.DisplayAdmin("views/customer/add.html")
	} else {
		// 新增或保存
		cus := make(map[string]interface{}, 0)
		cus["name"] = req.PostFormValue("name")
		cus["mobile"] = req.PostFormValue("mobile")
		cus["address"] = req.PostFormValue("address")
		cus["balance"] = req.PostFormValue("balance")
		cus["discount"] = req.PostFormValue("discount")
		// cus["company_id"] = 1
		// cus["is_delete"] = 0

		resMes := "添加"
		var err error
		var res int64
		if customerInfo != "" {
			// 有详情更新
			resMes = "编辑"
			res, err = customerModel.Update(cus, id)
		} else {
			// 无详情新增
			res, err = customerModel.Add(cus)
		}

		// 返回json值
		jr := &JsonResult{}
		if err != nil {
			jr.Code = 201
			jr.Message = resMes + "客户失败"
		} else {
			jr.Code = 200
			jr.Message = string(res)
		}
		r, _ := json.Marshal(jr)
		fmt.Fprintln(rep, string(r))
	}
}

// 顾客详情
func (c *CustomerController) Info() {
	// 根据id获取详情
	id := paramMap["id"]
	customerInfoMap := make(map[string]string, 0)
	customerInfoMap, err := customerModel.Info(id)
	jsonResult := &CustomerInfoJson{Code: 0}
	if err == nil {
		jsonResult.Code = 200
		jsonResult.Result = customerInfoMap
	}
	customerInfoByte, _ := json.Marshal(jsonResult)
	fmt.Fprintln(rep, string(customerInfoByte))
}
