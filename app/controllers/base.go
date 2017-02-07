package controllers

import (
	"fmt"
	"net/http"
	"reflect"
)

type BaseController struct {
	// req *http.Request
	// rep http.ResponseWriter
}

// 路由(/Index/index)
func (ba *BaseController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Base(w, r)
	// 解析路径
	// fmt.Fprintf(w, "%s \r\n", r.URL.Path)
	var funs = map[string]interface{}{
		"Index": &IndexController{req: r, rep: w},
	}

	con := "Index"
	ac := "Test"

	conVal := reflect.ValueOf(funs[con])
	method := conVal.MethodByName(ac)
	if method.IsValid() {
		method.Call([]reflect.Value{})
	}
}

func (ba *BaseController) Display() {
	fmt.Println("test")
}
