package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"go.app.dt.com/models"
	"go.app.dt.com/utils"
)

type OrderController struct {
	BaseController
}

var orderModel = &models.Order{}

// var productModel = &models.Product{}

type ProInfo struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	Count       string `json:"count"`
}

type CusInfo struct {
	CusID     string `json:"cus_id"`
	ExpCharge string `json:"exp_charge"`
}

type OrderInfo struct {
	ProInfos []ProInfo `json:"pro_infos"`
	CusInfo
}

// 新增订单
func (o *OrderController) Add() {
	// o.ALoginCheck()
	infoMap := make(map[string]interface{}, 0)
	infoMap["cus_id"] = "1"
	infoMap["exp_charge"] = "10"

	proInfoArr := make([]map[string]string, 0)

	proInfo := make(map[string]string, 0)
	proInfo["id"] = "1"
	proInfo["product_name"] = "test1"
	proInfo["price"] = "9.9"
	proInfo["count"] = "2"
	proInfoArr = append(proInfoArr, proInfo)

	proInfo = make(map[string]string, 0)
	proInfo["id"] = "2"
	proInfo["product_name"] = "test2"
	proInfo["price"] = "8.8"
	proInfo["count"] = "2"

	proInfoArr = append(proInfoArr, proInfo)
	infoMap["pro_infos"] = proInfoArr

	jsonByte, _ := json.Marshal(infoMap)

	if req.Method == "GET" {
		// 产品列表
		productList := productModel.List()
		listByte, _ := json.Marshal(productList)
		assign["List"] = string(listByte)
		o.DisplayAdmin("views/order/add.html")
	} else {
		var orderInfo OrderInfo
		infoJSON := string(jsonByte)
		err := json.Unmarshal([]byte(infoJSON), &orderInfo)
		if err == nil {
			strMaps := utils.StructArrToMapArr(orderInfo.ProInfos)
			lastID, err := orderModel.Add(orderInfo.CusID, orderInfo.ExpCharge, strMaps)
			fmt.Println(lastID, err)
		}
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
