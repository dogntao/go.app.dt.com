package controllers

import (
	"fmt"
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

//
type BaseController struct {
}

// 路由转发(/Index/index)
func (ba *BaseController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req = r
	rep = w

	con := "Index"
	ac := "Index"
	conVal := reflect.ValueOf(funs[con])
	method := conVal.MethodByName(ac)
	if method.IsValid() {
		method.Call([]reflect.Value{})
	}
}

func (ba *BaseController) Display() {
	fmt.Println("test")
}
