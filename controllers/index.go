package controllers

import (
	"fmt"
	"html/template"
	"strings"

	"go.app.dt.com/utils"
)

type User struct {
	userName string
	passWord string
}

type IndexController struct {
	BaseController
}

func (c *IndexController) Login() {
	user := &User{}
	if req.Method == "GET" {
		t, err := template.ParseFiles("views/index/index.html", "views/layouts/header.html", "views/layouts/footer.html")
		t.ExecuteTemplate(rep, "index", user)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		utils.Con()
		user.userName = req.PostFormValue("username")
		user.passWord = req.PostFormValue("password")
		fmt.Fprintln(rep, user)
	}
}

func (c *IndexController) Index() {
	// c.Display()
	for key, value := range strings.Split(req.RequestURI, "/") {
		fmt.Fprintln(rep, key, value)
	}

}
