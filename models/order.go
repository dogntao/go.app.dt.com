package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Order struct {
}

type OrderCount struct {
	QueryCount string
}

type OrderInfo struct {
	ID          string `json:"id"`
	OrderName   string `json:"order_name"`
	OrderDesc   string `json:"order_desc"`
	CusMobile   string `json:"cus_mobile"`
	CusName     string `json:"cus_name"`
	CusAddress  string `json:"cus_address"`
	CusDiscount string `json:"cus_discount"`
	ProCount    string `json:"pro_count"`
	ExpCharge   string `json:"exp_charge"`
	OriPrice    string `json:"ori_price"`
	Price       string `json:"price"`
	CreateDate  string `json:"create_date"`
	IsDelete    string `json:"is_delete"`
}

var orderTable = "cms_order"

// 新增订单
func (o *Order) Add(cusId, expCharge string, productCount map[string]interface{}) (lastId int64, err error) {
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
		// 批量查询产品详情
		productIds := []string{}
		productCounts := make([]interface{}, 0)
		for key, val := range productCount {
			productIds = append(productIds, key)
			productCounts = append(productCounts, val)
		}
		productModel := &Product{}
		productInfos := productModel.ListByIds(productIds)
		for k, v := range productInfos {
			// 生成总价格和总数
			orderInfoPrice, _ := strconv.ParseFloat(v["price"], 64)
			orderInfoCount, _ := strconv.ParseFloat(productCounts[k].(string), 64)
			oriPrice += orderInfoPrice * orderInfoCount
			proCount += orderInfoCount
			// 处理order_desc字段
			orderInfoNew = make(map[string]interface{}, 0)
			orderInfoNew["id"] = v["id"]
			orderInfoNew["name"] = v["product_name"]
			orderInfoNew["pirce"] = v["price"]
			orderInfoNew["count"] = orderInfoCount
			orderInfoNew["money"] = fmt.Sprintf("%.2f", orderInfoPrice*orderInfoCount)
			orderInfoNews = append(orderInfoNews, orderInfoNew)
		}
		orderInfoByte, _ := json.Marshal(orderInfoNews)
		data["order_desc"] = string(orderInfoByte)
		// 原价+快递费
		expChargeFloat, _ := strconv.ParseFloat(expCharge, 64)
		data["pro_count"] = proCount
		data["ori_price"] = fmt.Sprintf("%.2f", oriPrice+expChargeFloat)
		// 实际价格(客户折扣价格+快递费)
		cusDiscount, _ := strconv.ParseFloat(data["cus_discount"].(string), 64)
		cusDiscount = cusDiscount / 100
		data["price"] = fmt.Sprintf("%.2f", (oriPrice*cusDiscount)+expChargeFloat)
		data["create_date"] = time.Now().Unix()
		// fmt.Println(data)
		lastId, err = Dtsql.Insert(orderTable, data)
	}
	return
}

// 订单列表(返回总数和列表)
func (o *Order) Manage(seaStr string, pageIndex, pageSize int64) (total int, list []map[string]string) {
	con := ""
	bind := []string{}
	if seaStr != "" {
		con = "(id like ? OR order_name like ? OR cus_id like ? OR cus_mobile like ? OR cus_name like ?)"
		for i := 0; i < 5; i++ {
			seaStr = "%" + seaStr + "%"
			bind = append(bind, seaStr)
		}
	} else {
		// bind = append(bind, "")
	}
	// 查询总数
	var orderCount OrderCount
	err := Dtsql.Query(orderCount, orderTable, con, bind)
	err = Dtsql.FetchRow()
	total, _ = strconv.Atoi(Dtsql.RetMap["queryCount"])

	// 查询列表
	var orderInfo OrderInfo
	if con == "" {
		con = "1=1"
	}
	con = con + " LIMIT ?,?"
	bind = append(bind, fmt.Sprintf("%d", (pageIndex-1)*pageSize))
	bind = append(bind, fmt.Sprintf("%d", pageSize))
	err = Dtsql.Query(orderInfo, orderTable, con, bind)
	err = Dtsql.FetchAll()
	list = Dtsql.RetRows

	checkErr(err)
	return
}
