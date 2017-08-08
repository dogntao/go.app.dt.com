package controllers

import (
	"encoding/json"
	"fmt"

	"go.app.dt.com/models"
)

type OrderController struct {
	BaseController
}

var orderModel = &models.Order{}

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

func (o *OrderController) Add() {
	// o.ALoginCheck()
	infoMap := make(map[string]interface{}, 0)
	infoMap["cus_id"] = "dt"
	infoMap["exp_charge"] = "10"

	proInfoArr := make([]map[string]string, 0)
	proInfo := make(map[string]string, 0)
	proInfo["id"] = "1"
	proInfo["product_name"] = "test1"
	proInfo["price"] = "9.9"
	proInfo["count"] = "2"
	proInfoArr = append(proInfoArr, proInfo)
	proInfo["id"] = "2"
	proInfo["product_name"] = "test2"
	proInfo["price"] = "8.8"
	proInfo["count"] = "2"
	proInfoArr = append(proInfoArr, proInfo)
	infoMap["pro_infos"] = proInfoArr

	jsonByte, _ := json.Marshal(infoMap)
	fmt.Println(string(jsonByte))

	if req.Method == "GET" {
		var orderInfo OrderInfo
		infoJSON := string(jsonByte)
		err := json.Unmarshal([]byte(infoJSON), &orderInfo)
		if err == nil {
			// fmt.Println(orderInfo.CusID)
			// fmt.Println(orderInfo.ExpCharge)
			// fmt.Println(orderInfo.ProInfos)
			// lastID, err := orderModel.Add(orderInfo.CusID, orderInfo.ExpCharge, orderInfo.ProInfos)
		}
		// fmt.Println(strconv.FormatInt(time.Now().Unix(), 10))
		// o.DisplayAdmin("views/order/add.html")
	} else {
		// var orderInfo OrderInfo
		// infoJSON := req.PostFormValue("info")
		// err := json.Unmarshal([]byte(infoJSON), &orderInfo)
		// if err == nil {
		// 	lastId, err := orderModel.Add()
		// }
	}
}
