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
	Id           string `json:"id"`
	Product_name string `json:"product_name"`
	Price        string `json:"price"`
	Count        string `json:"count"`
	Purcase      string `json:"purcase"`
}

// 产品首页
func (p *ProductController) Index() {
	fmt.Println("index")
	p.Display("views/product/index.html")
}

// 产品
func (p *ProductController) Add() {
	if req.Method == "GET" {
		p.DisplayAdmin("views/product/add.html")
	} else {
		// 新增
		pro := make(map[string]interface{}, 0)
		pro["product_name"] = req.PostFormValue("product_name")
		pro["price"] = req.PostFormValue("price")
		pro["count"] = req.PostFormValue("count")
		pro["is_delete"] = 1
		if req.PostFormValue("is_delete") == "1" {
			pro["is_delete"] = 0
		}
		res, err := productModel.Add(pro)
		// 返回json值
		jr := &JsonResult{}
		if err != nil {
			jr.Code = 201
			jr.Message = "新增产品失败"
		} else {
			jr.Code = 200
			jr.Message = string(res)
		}
		r, _ := json.Marshal(jr)
		fmt.Fprintln(rep, string(r))
	}
}

// 产品列表
func (p *ProductController) List() {
	if req.Method == "GET" {
		// 状态
		isDelete := paramMap["is_delete"]
		if isDelete == "" {
			isDelete = "-1"
		}
		// 列表
		productList := productModel.List(isDelete)
		// 处理进货
		for k := range productList {
			productList[k]["purcase"] = ""
		}
		listByte, _ := json.Marshal(productList)
		assign["List"] = string(listByte)
		assign["isDelete"] = isDelete
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
		jr := &JsonResult{}
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

// 上下线
func (p *ProductController) ChangeStatus() {
	id := paramMap["id"]
	if req.Method == "POST" {
		// 新增或保存
		pro := make(map[string]interface{}, 0)
		pro["is_delete"] = req.PostFormValue("is_delete")
		res, err := productModel.Update(pro, id)
		// 返回json值
		jr := &JsonResult{}
		if err != nil {
			jr.Code = 201
			jr.Message = "产品上/下架失败"
		} else {
			jr.Code = 200
			jr.Message = string(res)
		}
		r, _ := json.Marshal(jr)
		fmt.Fprintln(rep, string(r))
	}
}
