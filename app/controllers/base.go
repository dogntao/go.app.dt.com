package controller

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type BaseController struct {
	req *http.Request
	rep http.ResponseWriter
}

// 路由(/Index/index)
func (ba *BaseController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Base(w, r)
	// 解析路径
	fmt.Fprintf(w, "%s \r\n", r.URL.Path)
	// 获取do
	do := r.FormValue("do")
	do = strings.Title(do)
	// fmt.Fprintf(w, "%s \r\n", do)
	//调用do方法
	controlelr := &BaseController{}
	controlelr.req = r
	controlelr.rep = w
	conVal := reflect.ValueOf(controlelr)
	method := conVal.MethodByName(do)
	if method.IsValid() {
		method.Call([]reflect.Value{})
	}
}

// func Base(w http.ResponseWriter, r *http.Request) {
// 	// 解析路径
// 	fmt.Fprintf(w, "%s \r\n", r.URL.Path)

// 	// 解析参数
// 	r.ParseForm()
// 	for k, v := range r.Form {
// 		fmt.Fprintf(w, "%s,%s \r\n", k, v[0])
// 	}

// 	// 获取do
// 	do := r.FormValue("do")
// 	do = strings.Title(do)
// 	// fmt.Fprintf(w, "%s \r\n", do)

// 	//调用do方法
// 	controlelr := &BaseController{}
// 	controlelr.req = r
// 	controlelr.rep = w
// 	conVal := reflect.ValueOf(controlelr)
// 	method := conVal.MethodByName(do)
// 	if method.IsValid() {
// 		method.Call([]reflect.Value{})
// 	}
// }
