package controller

import (
	"fmt"
)

func (c *Controller) Test() {
	c.req.ParseForm()
	for k, v := range c.req.Form {
		fmt.Printf("%s,%s \r\n", k, v[0])
	}
	fmt.Println("test controller")
}

func (c *Controller) Index() {
	fmt.Println("index controller")
}
