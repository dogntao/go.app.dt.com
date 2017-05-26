package controllers

import (
	"fmt"
	"html/template"
	"strings"

	"go.app.dt.com/models"
	"go.app.dt.com/utils"
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
		check := user.LoginCheck()
		if check {
			// 保存cookie
			delete(models.Dtsql.RetMap, "pass_word")
			utils.SetCooke(rep, "user_info", models.Dtsql.RetMap)
		}
	}
}

func (c *IndexController) Index() {
	// c.Display()
	for key, value := range strings.Split(req.RequestURI, "/") {
		fmt.Fprintln(rep, key, value)
	}

}
