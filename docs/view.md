### view require.js+layui+vue
#### 一、require.js
> `require.js` js文件的异步加载，管理模块之间的依赖性(http://www.ruanyifeng.com/blog/2012/11/require_js.html)

> `reuire-css.js` 按需加载样式文件(https://segmentfault.com/a/1190000002390643)

##### &emsp; 1.`config.js`(reuqire.js配置)示例
```
require.config({
    baseUrl: "/js/components",
    paths: {
        "jquery": "jquery.min",
        "vue": "vue",
        "vue1": "vue1",
        "layui": "/js/components/layui/layui",
        "select2": "/js/components/select2/select2.min",
        "layui-private": "/js/private/layui-private",
        "vue-private": "/js/private/vue-private"
    },
    map: {
        "*": {
            "css": "css.min",
            "layui": "layui-private",
            // "select2": "select2-private",
            "vue": "vue-private"
        },
        "layui-private": {
            "layui": "layui",
        },
        // "select2-private": {
        //     "select2": "select2",
        // },
        "vue-private": {
            "vue": "vue1",
        }
    },
    shim: {
        "layui": {
            deps: ["css!/js/components/layui/css/layui.css"],
            exports: "layui"
        },
        "select2": {
            deps: ["css!/js/components/select2/select2.min.css"],
            exports: "select2"
        },
    }
})
```
##### &emsp; 2.调用require.js
```
<script src="/js/components/require.js"></script>
<script src="/js/config.js"></script>
<!-- 立马加载layui,控制页面样式 -->
<script src="/js/components/layui/layui.js"></script>
<link rel="stylesheet" href="/js/components/layui/css/layui.css">
<link rel="stylesheet" href="/css/layout_admin.css">
<script>
    require(['jquery', 'vue', 'layui'], function ($, Vue) {
        layui.use(['element'], function () { });
        new Vue({
            el: '#left',
            data: {
                con: '{{.Con}}',
                act: '{{.Act}}',
            },
        })
    })
</script>
```

#### 二、`layui`
> `layui` 非常棒的前端框架(http://www.layui.com/)

> 修改layui的components的路径
```
define(['layui'], function(layui) {
    layui.config({'dir':"/js/components/layui/"});
    return layui;
})
```

#### 三、`vue`
> `vue` 构建用户界面的渐进式框架(通过控制数据来控制页面显示效果)(https://cn.vuejs.org/index.html)

> 由于vue的默认输出和go的输出都是`{{}}`,修改go的输出方式为${}
```
define(['vue'], function(vue) {
    vue.config.delimiters = ['${', '}'];
    return vue;
})
```

#### 四、`vue双向绑定示例`
##### 1、model从数据库查找数据
```
// 查询详情
func (c *Customer) Info(id string) (info map[string]string, err error) {
	var customerInfo CustomerInfo
	con := "id=?"
	bind := []string{id}
	err = Dtsql.Query(customerInfo, cusTable, con, bind)
	err = Dtsql.FetchRow()
	info = Dtsql.RetMap
	return
}
```
##### 2、controller传递数据
```
// 根据id获取详情
id, ok := paramMap["id"]
customerInfoMap := make(map[string]string, 0)
if ok {
    customerInfoMap, _ = customerModel.Info(id)
}
// 传递json值
customerInfoByte, _ := json.Marshal(customerInfoMap)
customerInfo := string(customerInfoByte)

if req.Method == "GET" {
    // 获取到id参数,编辑
    assign["Id"] = id
    assign["Info"] = customerInfo
    c.DisplayAdmin("views/customer/add.html")
}
```

##### 3、vue显示数据
> 通过 info: jQuery.parseJSON('{{.Info}}')把传递过来的Info解析给了vue data里的info

> 通过  v-model="info.name" 实现了数据的双向绑定
```
{{define "add"}} {{template "left" .}}
<div id="add_order" v-cloak>
    <div method="POST" class="layui-form" action="">
        <div class="layui-form-item">
            <label class="layui-form-label">姓名</label>
            <div class="layui-input-block">
                <input type="text" name="name" v-model="info.name" required  lay-verify="required" placeholder="请输入姓名" autocomplete="off" class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">手机</label>
            <div class="layui-input-block">
                <input type="text" name="mobile" v-model="info.mobile" required  lay-verify="required" placeholder="请输入手机" autocomplete="off" class="layui-input">
            </div>
        </div>
        <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">地址</label>
            <div class="layui-input-block">
            <textarea name="address" v-model="info.address" placeholder="请输入地址" class="layui-textarea"></textarea>
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">余额</label>
            <div class="layui-input-block">
                <input type="text" name="balance" v-model="info.balance" placeholder="请输入余额" autocomplete="off" class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">折扣</label>
            <div class="layui-input-block">
                <input type="text" name="discount" v-model="info.discount" value="100" placeholder="请输入折扣" autocomplete="off" class="layui-input">
            </div>
        </div>

        <div class="layui-form-item">
            <div class="layui-input-block">
                <button class="layui-btn" lay-submit lay-filter="customerAdd">立即提交</button>
            </div>
        </div>
    </div>
</div>
<script>
    require(['jquery', 'vue', 'layui'], function($,Vue) {
        layui.use(['form','layer'], function() {
            var form = layui.form;
            var layer = layui.layer;
            form.on('submit(customerAdd)', function(data) {
                $.post("/Customer/add?id={{.Id}}", data.field, function(result) {
                    if (result.Code == 200) {
                        location.href = "/Customer/Manage";
                    } else {
                        layer.msg(result.Message)
                    }
                },"json")
            })
        });
        new Vue({
            el: '#add_order',
            data: {
                info: jQuery.parseJSON('{{.Info}}'),
            },
            methods:{
                // 搜索
                search:function(){
                    location.href="/Customer/Manage?search="+this.searchText
                },
                // 编辑
            }
        })
    })
</script>
{{template "footer"}} {{end}}
```

#### 五、`vue循环示例`
##### 1、model从数据库查找数据
&emsp;数据库模糊查找数据，并分页,返回总数和数据列表
```
// 列表(返回总数和列表)
func (c *Customer) Manage(seaStr string, pageIndex, pageSize int64) (total int, list []map[string]string) {
	con := ""
	bind := []string{}
	if seaStr != "" {
		con = "(id like ? OR name like ? OR mobile like ? OR address like ?)"
		for i := 0; i < 4; i++ {
			seaStr = "%" + seaStr + "%"
			bind = append(bind, seaStr)
		}
	} else {
		// bind = append(bind, "")
	}
	// 查询总数
	var cusCount customerCount
	err := Dtsql.Query(cusCount, cusTable, con, bind)
	err = Dtsql.FetchRow()
	total, _ = strconv.Atoi(Dtsql.RetMap["queryCount"])

	// 查询列表
	var cusInfo CustomerInfo
	if con == "" {
		con = "1=1"
	}
	con = con + " LIMIT ?,?"
	bind = append(bind, fmt.Sprintf("%d", (pageIndex-1)*pageSize))
	bind = append(bind, fmt.Sprintf("%d", pageSize))
	err = Dtsql.Query(cusInfo, cusTable, con, bind)
	err = Dtsql.FetchAll()
	list = Dtsql.RetRows

	checkErr(err)
	return
}
```
##### 2、controlelr传递数据
&emsp; 获取serarch参数，通过customerModel.Manage方法查找数据并分页
```
// 列表
func (c *CustomerController) Manage() {
	if req.Method == "GET" {
		// 搜索数据
		search := paramMap["search"]
		// 处理分页
		var page int64
		page = 1
		getPage, ok := paramMap["page"]
		if ok {
			page, _ = strconv.ParseInt(getPage, 10, 64)
		}

		total, list := customerModel.Manage(search, page, 10)
		listByte, _ := json.Marshal(list)
		// 传递参数
		// assign["Total"] = math.Ceil(float64(total) / 10)
		assign["Total"] = total
		assign["List"] = string(listByte)
		assign["Search"] = search
		assign["Page"] = page
		c.DisplayAdmin("views/customer/manage.html")
	}
}
```
##### 3、vue 通过 v-for显示数据
&emsp; 通过`list: jQuery.parseJSON('{{.List}}')`将传递过来数据赋值给vue data里面的list

&emsp; 通过 `<tr v-for="item in list">` 遍历数据

&emsp; 通过 `${}`显示数据 <td>${item.id}</td>
```
{{define "manage"}} {{template "left" .}}
<div id="order_manage" v-cloak>
    <div style="width:400px; float:left; margin-bottom:10px">
        <input type="text" name="title" required lay-verify="required" placeholder="请输入id/姓名/手机/地址" autocomplete="off" class="layui-input"
            style="width:300px;float:left" v-model="searchText">
        <button class="layui-btn" style="float:left" @click="search()">查找</button>
    </div>
    <table class="layui-table">
        <!-- <colgroup>
            <col width="150">
            <col width="200">
            <col>
        </colgroup> -->
        <thead>
            <tr>
                <th>序号</th>
                <th>订单名称</th>
                <th>客户姓名</th>
                <th>客户手机号</th>
                <th>客户地址</th>
                <th>客户折扣</th>
                <th>产品总数</th>
                <th>快递费</th>
                <th>原价</th>
                <th>实际价格</th>
                <th>操作</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="item in list">
                <td>${item.id}</td>
                <td>${item.order_name}</td>
                <td>${item.cus_name}</td>
                <td>${item.cus_mobile}</td>
                <td>${item.cus_address}</td>
                <td>${item.cus_discount}</td>
                <td>${item.pro_count}</td>
                <td>${item.exp_charge}</td>
                <td>${item.ori_price}</td>
                <td>${item.price}</td>
                <td>
                    <a class="layui-btn layui-btn-normal" href="/Order/info?id=${item.id}">查看</a>
                    <a class="layui-btn layui-btn-normal" href="/Order/Delete?id=${item.id}">删除</a>
                </td>
            </tr>
        </tbody>
    </table>
    <div id="page"></div>
</div>
<script>
    require(['jquery', 'vue', 'layui'], function ($, Vue) {
        layui.use(['laypage'], function () {
            var laypage = layui.laypage
            // v2版本分页
            laypage.render({
                elem: 'page',
                count: '{{.Total}}',
                groups: 8, //连续显示分页数
                curr: '{{.Page}}',
                jump: function (obj, first) {
                    if (first != true) {
                        var search = '{{.SearchText}}';
                        location.href = "/Order/Manage?search=" + search + "&page=" +obj.curr;
                    }
                }
            });
        });
        new Vue({
            el: '#order_manage',
            data: {
                list: jQuery.parseJSON('{{.List}}'),
                searchText: '{{.Search}}',
            },
            methods: {
                // 搜索
                search: function () {
                    location.href = "/Order/Manage?search=" + this.searchText
                },
            }
        })
    })
</script>
{{template "footer"}} {{end}}
```