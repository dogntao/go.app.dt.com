package controller

import (
	"fmt"
)

func (c *Controller) Test() {
	fmt.Println("test controller")
}

func (c *Controller) Index() {
	fmt.Println("index controller")
}
