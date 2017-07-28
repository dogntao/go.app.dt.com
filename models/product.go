package models

import (
	"fmt"
)

type Product struct {
}

type ProductInfo struct {
	id           int64
	product_name string
	price        float64
	count        int64
}

var productTable = "cms_product"

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

// 批量更新
func (p *Product) UpdateProducts(upDatas []map[string]string) {
	affRow, err := Dtsql.UpdateMulti(productTable, upDatas, "id")
	fmt.Println(affRow, err)
	return
}
