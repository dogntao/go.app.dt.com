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
)

type Router struct {
	con string
	ac  string
}

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
	router := &Router{}
	for key, value := range strings.Split(req.RequestURI, "/") {
		if key == 1 {
			router.con = value
		} else if key == 2 {
			router.ac = value
		}
	}

	// 默认跳转到首页
	if router.con == "" {
		router.con = "Index"
	}
	if router.ac == "" {
		router.ac = "Index"
	}

	if funs[router.con] != nil {
		// 反射调用对应controller对应方法
		conVal := reflect.ValueOf(funs[router.con])
		method := conVal.MethodByName(router.ac)
		if method.IsValid() {
			method.Call([]reflect.Value{})
		}
	}
}
