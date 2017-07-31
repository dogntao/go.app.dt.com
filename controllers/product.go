package controllers

import (
	"encoding/json"
	"fmt"
	"reflect"

	"strings"

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
				upProduct = make(map[string]string, 0)
				inPurcase = make(map[string]string, 0)
				for i := 0; i < rT.NumField(); i++ {
					tName = rT.Field(i).Name
					tName = strings.ToLower(tName)
					tVal = rV.Field(i).Interface().(string)
					// 修改产品去掉进货
					if tName != "purcase" {
						upProduct[tName] = tVal
					}
					// 生成进货数据
					if tName == "id" || tName == "purcase" {
						inPurcase[tName] = tVal
						// 进货值为空删除
						if tName == "purcase" && (tVal == "" || tVal == "0") {
							inPurcase = make(map[string]string, 0)
						}
					}
				}
				upProducts = append(upProducts, upProduct)
				// 只生成有值的进货
				if len(inPurcase) > 0 {
					inPurcases = append(inPurcases, inPurcase)
				}
			}
		} else {
			fmt.Println(err)
		}
		// fmt.Println(upProducts)
		// fmt.Println(inPurcases)
		// upProductInfo.idArr =
		affRows, err := productModel.UpdateProducts(upProducts, inPurcases)
		fmt.Println(affRows, err)

		// 返回json值
		jr := &jsonResult{}
		if err != nil {
			jr.Code = 201
			jr.Message = "修改产品失败"
		} else {
			jr.Code = 200
			jr.Message = string(affRows)
		}
		r, _ := json.Marshal(jr)
		fmt.Fprintln(rep, string(r))
	}
}
