package models

import "encoding/json"
import "strings"

type Product struct {
}

type ProductInfo struct {
	ID          int64   `json:"id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Count       int64   `json:"count"`
	IsDelete    int64   `json:"is_delete"`
}

var productTable = "cms_product"
var purchaseTable = "cms_product_purchase"

// 新增产品
func (p *Product) Add(data map[string]interface{}) (lastId int64, err error) {
	lastId, err = Dtsql.Insert(productTable, data)
	return
}

// 产品列表
func (p *Product) List(isDelete string) (list []map[string]string) {
	var productInfo ProductInfo
	con := ""
	bind := []string{}
	if isDelete != "-1" && isDelete != "" {
		con = "is_delete=?"
		bind = append(bind, isDelete)
	}
	err := Dtsql.Query(productInfo, productTable, con, bind)
	err = Dtsql.FetchAll()
	checkErr(err)
	list = Dtsql.RetRows
	return
}

// 指定产品列表
func (p *Product) ListByIds(ids []string) (list []map[string]string) {
	var productInfo ProductInfo
	bind := []string{}
	conData := []string{}
	for _, val := range ids {
		bind = append(bind, val)
		conData = append(conData, "?")
	}
	con := "id in ("
	con += strings.Join(conData, ",")
	con += ")"
	err := Dtsql.Query(productInfo, productTable, con, bind)
	err = Dtsql.FetchAll()
	checkErr(err)
	list = Dtsql.RetRows
	return
}

// 批量更新(更新产品，新增库存)
func (p *Product) UpdateProducts(upDatas, inPurcases []map[string]string) (affRow int64, err error) {
	if len(upDatas) > 0 {
		// 修改产品
		// fmt.Println(upDatas)
		affRow, err = Dtsql.UpdateMulti(productTable, upDatas, "id")
		// 新增库存信息
		if len(inPurcases) > 0 {
			// fmt.Println(inPurcases)
			data := make(map[string]interface{}, 0)
			purcaseBype, _ := json.Marshal(inPurcases)
			data["pro_purchase"] = string(purcaseBype)
			data["is_delete"] = 0
			Dtsql.Insert(purchaseTable, data)
			// fmt.Println(lastId, err)
		}
	}
	return
}
