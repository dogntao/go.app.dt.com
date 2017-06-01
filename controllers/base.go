package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"go.app.dt.com/utils"
)

type BaseController struct {
}

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
	fmt.Println(pageName)
	tem.ExecuteTemplate(rep, pageName, "")
}

// 显示前台页面
func (b *BaseController) Display(page string) {
	tem, _ := template.ParseFiles(page, "views/layouts/index/left.html", "views/layouts/index/header.html", "views/layouts/index/footer.html")
	pageName := getFileName(page)
	fmt.Println(pageName)
	tem.ExecuteTemplate(rep, pageName, "")
}

// 后台登录验证(未登录跳转到登录页)
func (b *BaseController) ALoginCheck() {
	// 获取cookie
	_, err := utils.GetCookie(req, "user_info")
	if err != nil {
		http.Redirect(rep, req, "/Admin/Login", http.StatusFound)
	}
}
