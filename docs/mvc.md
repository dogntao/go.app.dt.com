### mvc
#### 一.controller
> 通过自定义路由，解析到对一个controller里面对应方法

#### 二.model
> 1.通过`import`方法引入model目录

> 2.初始化model

> 3.调用对应model对应方法
示例:
```
import (
	"go.app.dt.com/models"
)
var customerModel = &models.Customer{}
customerModel.Info(id)
```

#### 三.view
##### &emsp;1.加载模板
> a. go 通过`template.ParseFiles`把所有嵌套模板解析到模板里面，模板之间相互独立，并行存在，内部存储的是map关系(key是模板名称,value是模板内容)

> b. 然后通过`ExecuteTemplate`执行相应子模板内容
```
func (b *BaseController) DisplayAdmin(page string) {
	tem, _ := template.ParseFiles(page, "views/layouts/admin/left.html", "views/layouts/admin/header.html", "views/layouts/admin/footer.html")
	pageName := getFileName(page)
	assign["Con"] = con
	assign["Act"] = act
	tem.ExecuteTemplate(rep, pageName, assign)
}
```
##### &emsp;2.模板定义
> 在实际开发过程中经常要加载header，footer等

> 用define定义模板:{{define "header"}},用end结束定义:{{end}}
##### &emsp;3.模板嵌套模板
> 1.引入header {{define "add"}} {{template "header" .}}

> 2.引入footer {{template "footer"}} {{end}}

> `注:{{template "header" .}} .用于多页面参数打通`

##### &emsp;3.模板显示数据
> 1.controller传递参数 

> var assign = make(map[string]interface{})
```
ExecuteTemplate(req,pageName,assign),第三个参数`assign`即为往模板页面传递的数据
```

> 2.view 显示参数

> go通过`{{}}`来显示传递过来的数据，`{{.}}`表示当前对象,字段必须`首字母大写`
```
{{.FileName}}
```

