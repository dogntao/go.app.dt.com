package controllers

import (
	"net/http"
	"reflect"
	"strings"
)

// 保存输入和输出
var (
	req      *http.Request
	rep      http.ResponseWriter
	con      string
	act      string
	paramMap map[string]string
)

// 注册路由
var funs = map[string]interface{}{
	"Index":    &IndexController{},
	"Course":   &CourseController{},
	"Student":  &StudentController{},
	"Product":  &ProductController{},
	"About":    &AboutController{},
	"Admin":    &AdminController{},
	"Order":    &OrderController{},
	"Customer": &CustomerController{},
}

// 路由转发(/Index/index)
func IndexRouter(w http.ResponseWriter, r *http.Request) {
	// 全局输入和输出
	req = r
	rep = w

	// 排除/favicon.ico
	if req.RequestURI != "/favicon.ico" {
		// 区分地址和参数
		urlInfos := strings.Split(req.RequestURI, "?")
		path := urlInfos[0]
		// 根据/解析对应controller对应方法
		for key, value := range strings.Split(path, "/") {
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

		// 解析参数
		params := req.URL.Query()
		paramMap = make(map[string]string)
		for k, v := range params {
			if len(v) > 0 && v[0] != "" {
				paramMap[k] = v[0]
			}
		}
		// fmt.Println(paramMap)

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
}
