package controllers

import (
	"fmt"
	"html/template"
	"strings"

	"go.app.dt.com/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Login() {
	if req.Method == "GET" {
		t, err := template.ParseFiles("views/index/index.html", "views/layouts/header.html", "views/layouts/footer.html")
		t.ExecuteTemplate(rep, "index", "")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		user := &models.UserInfo{}
		user.UserName = req.PostFormValue("username")
		user.PassWord = req.PostFormValue("password")
		fmt.Fprintln(rep, user)

		user.LoginCheck()
	}
}

func (c *IndexController) Index() {
	// c.Display()
	for key, value := range strings.Split(req.RequestURI, "/") {
		fmt.Fprintln(rep, key, value)
	}

}
