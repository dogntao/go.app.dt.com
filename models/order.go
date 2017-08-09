package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Order struct {
}

var orderTable = "cms_order"

// 新增订单
func (o *Order) Add(cusId, expCharge string, orderInfo []map[string]interface{}) (lastId int64, err error) {
	data := make(map[string]interface{}, 0)
	// 获取用户信息
	cusModel := &Customer{}
	lastId = 0
	cusInfo, err := cusModel.Info(cusId)
	if err == nil {
		data["order_name"] = cusInfo["name"] + "-订单-" + time.Now().Format("2006-01-02 15:04")
		data["cus_id"] = cusId
		data["cus_mobile"] = cusInfo["mobile"]
		data["cus_name"] = cusInfo["name"]
		data["cus_address"] = cusInfo["address"]
		data["cus_discount"] = cusInfo["discount"]
		data["exp_charge"] = expCharge
		// 获取产品总数和价格
		var oriPrice, proCount float64
		oriPrice, proCount = 0, 0
		// 处理数据
		orderInfoNews := make([]map[string]interface{}, 0)
		orderInfoNew := make(map[string]interface{}, 0)
		for _, val := range orderInfo {
			// 生成总价格和总数
			orderInfoPrice, _ := strconv.ParseFloat(val["price"].(string), 64)
			orderInfoCount, _ := strconv.ParseFloat(val["count"].(string), 64)
			oriPrice += orderInfoPrice * orderInfoCount
			proCount += orderInfoCount
			// 处理order_desc字段
			orderInfoNew = make(map[string]interface{}, 0)
			orderInfoNew["id"] = val["id"]
			orderInfoNew["name"] = val["product_name"]
			orderInfoNew["pirce"] = val["price"]
			orderInfoNew["count"] = val["count"]
			orderInfoNew["money"] = fmt.Sprintf("%.2f", orderInfoPrice*orderInfoCount)
			orderInfoNews = append(orderInfoNews, orderInfoNew)
		}
		orderInfoByte, _ := json.Marshal(orderInfoNews)
		data["order_desc"] = string(orderInfoByte)
		// 快递费
		expChargeFloat, _ := strconv.ParseFloat(expCharge, 64)
		data["pro_count"] = proCount
		data["ori_price"] = fmt.Sprintf("%.2f", oriPrice+expChargeFloat)
		// 客户折扣
		cusDiscount, _ := strconv.ParseFloat(data["cus_discount"].(string), 64)
		cusDiscount = cusDiscount / 100
		data["price"] = fmt.Sprintf("%.2f", (oriPrice*cusDiscount)+expChargeFloat)
		data["create_date"] = time.Now().Unix()

		lastId, err = Dtsql.Insert(orderTable, data)
	}
	return
}

// 订单列表
func (o *Order) List() (list []map[string]string) {
	var productInfo ProductInfo
	con := "is_delete=?"
	bind := []string{"0"}
	err := Dtsql.Query(productInfo, productTable, con, bind)
	err = Dtsql.FetchAll()
	checkErr(err)
	list = Dtsql.RetRows
	return
}
