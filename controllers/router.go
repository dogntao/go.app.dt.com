package controllers

import (
	"net/http"
	"reflect"
	"strings"
)

// 保存输入和输出
var (
	req *http.Request
	rep http.ResponseWriter
	con string
	act string
)

// 注册路由
var funs = map[string]interface{}{
	"Index":    &IndexController{},
	"Admin":    &AdminController{},
	"Order":    &OrderController{},
	"Customer": &CustomerController{},
}

// 路由转发(/Index/index)
func IndexRouter(w http.ResponseWriter, r *http.Request) {
	// 全局输入和输出
	req = r
	rep = w

	// 根据/解析对应controller对应方法
	for key, value := range strings.Split(req.RequestURI, "/") {
		// 首字母大写
		value = strings.Title(value)
		if key == 1 {
			con = value
		} else if key == 2 {
			act = value
		}
	}

	// 默认跳转到ndex方法
	if act == "" {
		act = "Index"
	}

	if funs[con] != nil {
		// 反射调用对应controller对应方法
		conVal := reflect.ValueOf(funs[con])
		method := conVal.MethodByName(act)
		if method.IsValid() {
			method.Call([]reflect.Value{})
		}
	} else {
		// 页面没找到跳转到首页
		http.Redirect(rep, req, "/index/index", http.StatusFound)
	}
}
