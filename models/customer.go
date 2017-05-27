package models

type Customer struct {
}

var cusTable = "cms_customer"

// 新增
func (c *Customer) Add(data map[string]interface{}) (lastId int64, err error) {
	lastId, err = Dtsql.Insert(cusTable, data)
	return
}

// 更新
func (c *Customer) Update(upData map[string]interface{}, conStr string) (affRow int64, err error) {
	affRow, err = Dtsql.Update(cusTable, upData, conStr)
	return
}

// 列表
// func (c *Customer)  {

// }
