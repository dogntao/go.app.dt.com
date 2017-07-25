package controllers

import "go.app.dt.com/models"

type ProductController struct {
	BaseController
}

var productModel = &models.Product{}

// 产品
func (p *ProductController) Index() {
	p.Display("views/product/product.html")
}

// 产品
func (p *ProductController) Update() {
	productModel.UpdateProducts()
}
