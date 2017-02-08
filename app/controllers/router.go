package controllers

import (
	"net/http"
	"reflect"
)

// 保存输入和输出
var (
	req *http.Request
	rep http.ResponseWriter
)

// 注册路由
var funs = map[string]interface{}{
	"Index": &IndexController{},
}

// 路由
type Router struct {
}

// 路由转发(/Index/index)
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req = r
	rep = w

	con := "Index"
	ac := "Test"
	conVal := reflect.ValueOf(funs[con])
	method := conVal.MethodByName(ac)
	if method.IsValid() {
		method.Call([]reflect.Value{})
	}
}
