package models

type Product struct {
}

var productTable = "cms_product"

// 批量更新
func (p *Product) UpdateProducts() {
	upDatas := make([]map[string]string, 5)
	data := make(map[string]string)

	data["id"] = "1"
	data["name"] = "name_1"
	data["title"] = "title_1"
	upDatas[0] = data

	data = make(map[string]string)
	data["id"] = "2"
	data["name"] = "name_2"
	data["title"] = "title_2"
	upDatas[1] = data

	Dtsql.UpdateMulti(productTable, upDatas, "id")
	return
}
