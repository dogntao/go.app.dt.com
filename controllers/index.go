package controllers

import (
	"fmt"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	fmt.Println(act)
	c.Display("views/index/index.html")
}

func (c *IndexController) Course() {
	// c.Display()
	fmt.Println(act)
	c.Display("views/index/course.html")
}
