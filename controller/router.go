package controller

import (
	"fmt"
	"net/http"
	"reflect"
)

type Controller struct {
}

func Base(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//	for k, v := range r.Form {
	//		fmt.Fprintf(w, "%s,%s \r\n", k, v[0])
	//	}

	//获取do
	do := r.FormValue("do")
	if do == "" {
		do = "Index"
	}
	do
	fmt.Fprintf(w, "%s \r\n", do)

	controlelr := &Controller{}
	conVal := reflect.ValueOf(controlelr)
	//	method := conVal.MethodByName("")
	controlelr.Test()
	fmt.Fprintf(w, "%s,\r\n", conVal)
}
