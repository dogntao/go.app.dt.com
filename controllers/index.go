package controllers

type IndexController struct {
	BaseController
}

// 首页
func (i *IndexController) Index() {
	i.Display("views/index/index.html")
}
