package controllers

import (
	"net/http"
	"reflect"
	"strings"

	"fmt"

	"go.app.dt.com/utils"
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
	"Index": &IndexController{},
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

	if funs[router.con] != nil {
		// 获取cookie
		cookieMap, err := utils.GetCookie(req, "user_info")
		fmt.Println(cookieMap, err)
		if err != nil {
			// 无cookie跳转到登录页
			conVal := reflect.ValueOf(funs["Index"])
			method := conVal.MethodByName("Login")
			if method.IsValid() {
				method.Call([]reflect.Value{})
			}
		} else {
			// 有cookie跳转到对应页面
			fmt.Println(cookieMap)
			// 反射调用对应controller对应方法
			conVal := reflect.ValueOf(funs[router.con])
			method := conVal.MethodByName(router.ac)
			if method.IsValid() {
				method.Call([]reflect.Value{})
			}
		}

	}
}
