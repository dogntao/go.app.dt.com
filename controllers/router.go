package controllers

import (
	"encoding/json"
	"fmt"
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
	// 调试输出
	if funs[router.con] != nil {
		// 反射调用对应controller对应方法
		conVal := reflect.ValueOf(funs[router.con])
		method := conVal.MethodByName(router.ac)

		// 获取cookie
		cookie, _ := req.Cookie("user_info")
		if cookie != nil {
			cookieValue := cookie.Value
			cookieMap := make(map[string]interface{})
			err := json.Unmarshal([]byte(cookieValue), &cookieMap)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(cookieMap)
		}

		if method.IsValid() {
			method.Call([]reflect.Value{})
		}
	}
}
