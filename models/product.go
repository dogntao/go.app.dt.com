package models

import (
	"fmt"
)

type Product struct {
}

var productTable = "test"

// 批量更新
func (p *Product) UpdateProducts() {
	upDatas := make([]map[string]string, 5)
	data := make(map[string]string)

	data["id"] = "1"
	data["name"] = "name1"
	data["title"] = "title1"
	upDatas[0] = data

	data = make(map[string]string)
	data["id"] = "2"
	data["name"] = "name2"
	data["title"] = "title2"
	upDatas[1] = data

	affRow, err := Dtsql.UpdateMulti(productTable, upDatas, "id")
	fmt.Println(affRow, err)
	return
}
