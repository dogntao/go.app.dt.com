### 路由
1. 把`静态文件`解析到`指定目录`
2. 把`根目录`解析到`自定义方法`
3. `监听`指定端口,创建`server`服务

&emsp;示例:
```
func main() {
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/image/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", con.IndexRouter)
	http.ListenAndServe(":6688", nil)
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
