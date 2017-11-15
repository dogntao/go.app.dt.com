### 自定义路由
```
http://127.0.0.1:6688/Product/List?is_delete=0
```
1. `?`后是参数
2. 用`/`区分controller和action
3. `用reflect`调用对应controller对应方法

&emsp;示例:
```
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
		// 默认跳转到index方法
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
```
#### 1.静态文件
> 利用`http.Handle`将静态文件夹转发到项目对应路径 
``` 
http.Handle("/js/", http.FileServer(http.Dir("public")))
```

&emsp;`http.Dir`返回文件夹相对路径地址

&emsp;`http.FileServer`获取到项目根目录handler

&emsp;这段话相当于把`/js/`路径转发到了项目根目录`/public/js/`下

#### 2.动态路由
> 利用`http.HandleFunc`将根目录转发到自定义路由方法 
``` 
http.HandleFunc("/", con.IndexRouter)
```

&emsp;`con.IndexRouter`自定义路由方法
