### mvc
#### 1.controller
> 通过自定义路由，解析到对一个controller里面对应方法

#### 2.model
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

#### 3.view
```
> 1.go 通过`template.ParseFiles`把所有嵌套模板解析到模板里面，模板之间相互独立，并行存在，内部存储的是map关系(key是模板名称,value是模板内容),
```