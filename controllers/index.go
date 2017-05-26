package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"go.app.dt.com/models"
	"go.app.dt.com/utils"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Login() {
	// 获取cookie
	_, err := utils.GetCookie(req, "user_info")
	if req.Method == "GET" {

		if err != nil {
			// 无cookie登录
			c.Display("views/index/index.html")
		} else {
			// 有cookie跳到新增订单页
			http.Redirect(rep, req, "http://127.0.0.1:6688/Order/Add", http.StatusMovedPermanently)
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
			fmt.Fprintln(rep, true)
		}
	}
}

func (c *IndexController) Index() {
	// c.Display()
	for key, value := range strings.Split(req.RequestURI, "/") {
		fmt.Fprintln(rep, key, value)
	}
}
