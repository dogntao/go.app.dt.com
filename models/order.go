package models

import "fmt"

type Order struct {
}

var orderTable = "cms_order"

// 新增订单
func (o *Order) Add(cusId, expCharge string, data []map[string]interface{}) {
	fmt.Println(cusId)
	fmt.Println(expCharge)
	fmt.Println(data)

	// cusModel := &Customer{}
	// cusInfo := cusModel.Info(cusId)

	// order := make(map[string]interface{}, 0)
	// order["order_desc"] = string(orderDesc)
	// order["cus_id"] = orderInfo.CusID
	// order["com_id"] = orderInfo.cus_id
	// order["pro_count"] = orderInfo.cus_id
	// order["ori_price"] = orderInfo.cus_id
	// order["price"] = orderInfo.cus_id
	// order["create_date"] = strconv.FormatInt(time.Now().Unix(), 10)
	// order["is_delete"] = "0"

	// lastId, err = Dtsql.Insert(orderTable, data)
	// return
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
