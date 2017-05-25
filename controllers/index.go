package controllers

import (
	"fmt"
	"html/template"
	"strings"

	"encoding/json"

	"net/http"

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
		check := user.LoginCheck()
		if check {
			// 保存cookie
			delete(models.Dtsql.RetMap, "pass_word")
			str, _ := json.Marshal(models.Dtsql.RetMap)
			cookie := &http.Cookie{Name: "user_info", Value: fmt.Sprintf("%s", string(str)), Path: "/", MaxAge: 86400}
			http.SetCookie(rep, cookie)
			fmt.Println(string(str))
		}
	}
}

func (c *IndexController) Index() {
	// c.Display()
	for key, value := range strings.Split(req.RequestURI, "/") {
		fmt.Fprintln(rep, key, value)
	}

}
