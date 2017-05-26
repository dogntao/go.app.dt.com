package controllers

import (
	"fmt"
	"html/template"
	"strings"
)

//
type BaseController struct {
}

// 根据路径获取文件名(不带后缀)
func getFileName(filePath string) string {
	pageArr := strings.Split(filePath, "/")
	pageName := pageArr[len(pageArr)-1]
	pageNameArr := strings.Split(pageName, ".")
	return pageNameArr[0]
}

// 显示页面
func (ba *BaseController) Display(page string) {
	tem, _ := template.ParseFiles(page, "views/layouts/left.html", "views/layouts/header.html", "views/layouts/footer.html")
	pageName := getFileName(page)
	fmt.Println(pageName)
	tem.ExecuteTemplate(rep, pageName, "")
}
