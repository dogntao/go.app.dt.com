package controller

import (
	"fmt"
	"text/template"
)

type test struct {
	Name string
}

func (c *Controller) Test() {
	//	c.req.ParseForm()
	//	for k, v := range c.req.Form {
	//		fmt.Printf("%s,%s \r\n", k, v[0])
	//	}
	//	fmt.Println("test controller")
	testData := &test{"dongtao"}
	//	testData := "dongtao"
	t, err := template.ParseFiles("view/index/index.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(c.rep, testData)
}

func (c *Controller) Index() {
	fmt.Println("index controller")
}
