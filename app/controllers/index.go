package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name  string
	Hobby [3]string
}

type IndexController struct {
	req *http.Request
	rep http.ResponseWriter
	BaseController
}

func (c *IndexController) Test() {
	c.req.ParseForm()
	for k, v := range c.req.Form {
		fmt.Printf("%s,%s \r\n", k, v[0])
	}
	Users := make([]User, 3)
	Users[0].Name = "dongtao"
	Users[0].Hobby[0] = "swimming"
	Users[0].Hobby[1] = "basketball"
	Users[1].Name = "dt"
	Users[1].Hobby[0] = "football"
	Users[1].Hobby[1] = "pingpang"

	t, err := template.ParseFiles("view/layouts/header.html", "view/index/index.html", "view/layouts/footer.html")
	t.ExecuteTemplate(c.rep, "index", Users)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *IndexController) Index() {
	c.Display()
	// fmt.Println("index controller")
}
