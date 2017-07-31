package controllers

import (
	"encoding/json"
	"fmt"
	"reflect"

	"go.app.dt.com/models"
)

type ProductController struct {
	BaseController
}

var productModel = &models.Product{}

type UpProductInfo struct {
	Id           string
	Product_name string
	Price        string
	Count        string
	Purcase      string
}

// 产品
func (p *ProductController) Index() {
	p.Display("views/product/product.html")
}

// 产品列表
func (p *ProductController) List() {
	if req.Method == "GET" {
		// 列表
		productList := productModel.List()
		// 处理进货
		for k, _ := range productList {
			productList[k]["purcase"] = ""
		}
		listByte, _ := json.Marshal(productList)
		assign["List"] = string(listByte)
		p.DisplayAdmin("views/product/list.html")
	}
}

// 产品更新
func (p *ProductController) Update() {
	if req.Method == "POST" {
		var listUpProductInfo []UpProductInfo
		listJson := req.PostFormValue("list")
		err := json.Unmarshal([]byte(listJson), &listUpProductInfo)
		upProducts := make([]map[string]string, 0)
		upProduct := make(map[string]string, 0)
		inPurcases := make([]map[string]string, 0)
		inPurcase := make(map[string]string, 0)
		tName := ""
		tVal := ""
		if err == nil {
			for _, val := range listUpProductInfo {
				rT := reflect.TypeOf(val)
				rV := reflect.ValueOf(val)
				upProducts = make(map[string]string, 0)
				for i := 0; i < rT.NumField(); i++ {
					tName = rT.Field(i).Name
					tVal = rV.Field(i).Interface()
					if tName != "Purcase" {
						upProduct[tName] = tVal
					} else {
						inPurca[tName] = tVal
					}
				}
				upProducts = append(upProducts, upProduct)
				inPurcases = append(inPurcases, inPurcase)
			}
		} else {
			fmt.Println(err)
		}
		fmt.Println(upProducts)
		fmt.Println(inPurcases)
		// upProductInfo.idArr =
		// productModel.UpdateProducts(listUpProductInfo)
	}
}
