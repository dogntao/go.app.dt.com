package controllers

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	// c.Display()
	c.Display("views/index/index.html")
}

func (c *IndexController) Course() {
	// c.Display()
	c.Display("views/index/course.html")
}
