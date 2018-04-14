### 1、动态路由规则:通过斜杠区分(例:course/index)调用对应struce对应方法
> 实现方法：将所有struct实例放到map里，通过(course)字符串找到对应struct实例，然后利用反射MethodByName和Call动态调用方法(index)
```
    // 反射调用对应controller对应方法
    conVal := reflect.ValueOf(funs[con])
    method := conVal.MethodByName(act)
    if method.IsValid() {
        method.Call([]reflect.Value{})
    }
```
### 2、数据库操作
> 通过github.com/go-sql-driver/mysql封装查询、插入/批量插入、更新/批量更新等方法
### 3、页面解析
> 通过 template.ParseFiles和ExecuteTemplate传递参数和解析页面
```
// 显示前台页面
func (b *BaseController) Display(page string) {
	tem, _ := template.ParseFiles(page, "views/layouts/index/about_left.html", "views/layouts/index/have_left.html", "views/layouts/index/no_left.html", "views/layouts/index/header.html", "views/layouts/index/footer.html")
	pageName := getFileName(page)
	assign["Con"] = con
	assign["Act"] = act
	tem.ExecuteTemplate(rep, pageName, assign)
}
```
### 4、html
> 使用require.js+vue+div+css+layui