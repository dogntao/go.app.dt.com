package controllers

import (
	"fmt"
	"net/http"

	"encoding/json"

	"go.app.dt.com/models"
	"go.app.dt.com/utils"
)

type AdminController struct {
	BaseController
}

// 登录
func (c *AdminController) Login() {
	// 获取cookie
	_, err := utils.GetCookie(req, "user_info")
	if req.Method == "GET" {
		// lr := &loginResult{}
		// lr.Code = 201
		// lr.Result = "faile"
		// r, _ := json.Marshal(lr)
		// fmt.Fprintln(rep, string(r))

		if err != nil {
			// 无cookie登录
			c.DisplayAdmin("views/admin/login.html")
		} else {
			// 有cookie跳到新增订单页
			http.Redirect(rep, req, "/Order/Add", http.StatusFound)
		}

	} else {
		user := &models.UserInfo{}
		user.UserName = req.PostFormValue("username")
		user.PassWord = req.PostFormValue("password")
		check := user.LoginCheck()
		lr := &JsonResult{}
		lr.Code = 201
		lr.Message = "用户名或密码错误"
		if check {
			// 保存cookie
			delete(models.Dtsql.RetMap, "pass_word")
			utils.SetCooke(rep, "user_info", models.Dtsql.RetMap)
			lr.Code = 200
			lr.Message = "success"
		}
		r, _ := json.Marshal(lr)
		fmt.Fprintln(rep, string(r))
	}
}

// 登出
func (c *AdminController) Logout() {
	// 删除cookie
	utils.DelCookie(rep, "user_info")
	http.Redirect(rep, req, "/Admin/Login", http.StatusFound)
}

func (c *AdminController) Index() {
	// c.Display()
	fmt.Println("1234")
}
