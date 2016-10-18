package controller

import (
	// "fmt"
	"net/http"
	"reflect"
	"strings"
)

type Controller struct {
	req *http.Request
	rep http.ResponseWriter
}

func Base(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// for k, v := range r.Form {
	// 	fmt.Fprintf(w, "%s,%s \r\n", k, v[0])
	// }

	//获取do
	do := r.FormValue("do")
	do = strings.Title(do)
	// fmt.Fprintf(w, "%s \r\n", do)

	//调用do方法
	controlelr := &Controller{}
	controlelr.req = r
	controlelr.rep = w
	conVal := reflect.ValueOf(controlelr)
	method := conVal.MethodByName(do)
	if method.IsValid() {
		method.Call([]reflect.Value{})
	}
}
