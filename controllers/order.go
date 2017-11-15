package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

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
func (o *OrderController) Manage() {
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

		total, list := orderModel.Manage(search, page, 10)
		listByte, _ := json.Marshal(list)
		// 传递参数
		assign["Total"] = total
		assign["List"] = string(listByte)
		assign["Search"] = search
		assign["Page"] = page
		o.DisplayAdmin("views/order/manage.html")
	}
}

// 订单详情
func (o *OrderController) Info() {
	// 根据id获取详情
	id := paramMap["id"]
	orderInfoMap := make(map[string]string, 0)
	orderInfoMap, _ = orderModel.Info(id)
	timeInt, _ := strconv.ParseInt(orderInfoMap["create_date"], 10, 64)
	tm := time.Unix(timeInt, 0)
	orderInfoMap["create_date_formate"] = tm.Format("2006-01-02 15:04")
	orderInfoByte, _ := json.Marshal(orderInfoMap)
	assign["OrderInfo"] = string(orderInfoByte)
	assign["OrderDesc"] = orderInfoMap["order_desc"]
	// fmt.Println(assign["Orderinfo"])
	o.DisplayAdmin("views/order/info.html")
}
