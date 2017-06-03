package controllers

type ProductController struct {
	BaseController
}

// 产品
func (p *ProductController) Index() {
	p.Display("views/product/product.html")
}
