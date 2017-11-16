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
#### 1.获取uri用于解析controller、action、参数
> 利用`req.RequestURI`获取uri地址
``` 
返回:/Product/List?is_delete=1
```
#### 2.解析区分controller、action/参数
> 利用`strings.Split()`切分uri地址
> `?`切分uri地址
``` 
返回[/Product/List is_delete=1]数组
第0个是contreoller+action
第1个是参数
```
#### 3.解析controller、action
> 利用`/`切分数组第0个
> 利用`strings.Title()`,将首字母大写，兼容大小写输入
``` 
返回[ Product List]数组
第1个是contreoller
第2个是action
```
#### 4.解析参数
> 利用`req.URL.Query()`解析参数
``` 
返回map[is_delete:[1]]
key是参数名称
value是参数值数组
```
#### 5.利用反射调用对应controller对应方法
> 利用`conVal := reflect.ValueOf(funs[con])`反射
> 利用`method := conVal.MethodByName(act)`查找方法
> 利用`method.Call([]reflect.Value{})`调用方法
``` 
返回map[is_delete:[1]]
key是参数名称
value是参数值
```
