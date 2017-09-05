package models

import "encoding/json"

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
func (p *Product) List() (list []map[string]string) {
	var productInfo ProductInfo
	con := "is_delete=?"
	bind := []string{"0"}
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
