package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"go.app.dt.com/models"
)

type OrderController struct {
	BaseController
}

var orderModel = &models.Order{}

// 新增订单
func (o *OrderController) Add() {
	// o.ALoginCheck()
	if req.Method == "GET" {
		id := paramMap["id"]
		// 产品列表
		productList := productModel.List("0")
		listByte, _ := json.Marshal(productList)
		assign["List"] = string(listByte)
		// 获取到id参数,编辑
		assign["Id"] = id
		o.DisplayAdmin("views/order/add.html")
	} else {
		// 用户
		userID := req.PostFormValue("user_id")
		// 快递费
		expCharge := req.PostFormValue("exp_charge")
		// 产品及个数
		count := req.PostFormValue("count")
		var productCount map[string]interface{}
		json.Unmarshal([]byte(count), &productCount)
		// 新增
		lastID, err := orderModel.Add(userID, expCharge, productCount)
		fmt.Println(lastID, err)
		// 返回json值
		jr := &JsonResult{}
		if err != nil {
			jr.Code = 201
			jr.Message = "新增订单失败"
		} else {
			jr.Code = 200
			jr.Message = string(lastID)
		}
		r, _ := json.Marshal(jr)
		fmt.Fprintln(rep, string(r))
	}
}

// 订单列表
func (o *OrderController) List() {
	if req.Method == "GET" {
		// 搜索参数
		search := paramMap["search"]
		// 分页
		page := 1
		getPage, ok := paramMap["page"]
		if ok {
			page, _ = strconv.Atoi(getPage)
			if page < 1 {
				page = 1
			}
		}
		total, list := orderModel.List(search, page, 10)
		fmt.Println(total)
		fmt.Println(list)
	}

}
