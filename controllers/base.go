package controllers

import (
	"html/template"
	"net/http"
	"strings"

	"go.app.dt.com/utils"
)

type BaseController struct {
}

var assign = make(map[string]interface{})

// 根据路径获取文件名(不带后缀)
func getFileName(filePath string) string {
	pageArr := strings.Split(filePath, "/")
	pageName := pageArr[len(pageArr)-1]
	pageNameArr := strings.Split(pageName, ".")
	return pageNameArr[0]
}

// 显示后台页面
func (b *BaseController) DisplayAdmin(page string) {
	tem, _ := template.ParseFiles(page, "views/layouts/admin/left.html", "views/layouts/admin/header.html", "views/layouts/admin/footer.html")
	pageName := getFileName(page)
	tem.ExecuteTemplate(rep, pageName, "")
}

// 显示前台页面
func (b *BaseController) Display(page string) {
	tem, _ := template.ParseFiles(page, "views/layouts/index/about_left.html", "views/layouts/index/have_left.html", "views/layouts/index/no_left.html", "views/layouts/index/header.html", "views/layouts/index/footer.html")
	pageName := getFileName(page)
	assign["Con"] = con
	assign["Act"] = act
	tem.ExecuteTemplate(rep, pageName, assign)
}

// 后台登录验证(未登录跳转到登录页)
func (b *BaseController) ALoginCheck() {
	// 获取cookie
	_, err := utils.GetCookie(req, "user_info")
	if err != nil {
		http.Redirect(rep, req, "/Admin/Login", http.StatusFound)
	}
}
