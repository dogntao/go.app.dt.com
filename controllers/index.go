package controllers

import (
	"fmt"
	"html/template"
	"strings"

	"go.app.dt.com/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Login() {
	if req.Method == "GET" {
		t, err := template.ParseFiles("views/index/index.html", "views/layouts/header.html", "views/layouts/footer.html")
		t.ExecuteTemplate(rep, "index", "")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		user := &models.UserInfo{}
		user.UserName = req.PostFormValue("username")
		user.PassWord = req.PostFormValue("password")
		fmt.Fprintln(rep, user)

		user.LoginCheck()

		// db := dbStore.GetConn()
		// defer dbStore.RetConn(db)
		// fmt.Println(db)
		// fmt.Println("con")

		// 插入数据
		// stmt, err := db.Prepare("INSERT INTO cms_company (user_name,pass_word) values(?,?)")
		// checkErr(err)
		// res, err := stmt.Exec("test1", "123456")
		// checkErr(err)
		// id, err := res.LastInsertId()
		// checkErr(err)
		// fmt.Println(id)

		// user.userName = req.PostFormValue("username")
		// user.passWord = req.PostFormValue("password")
		// fmt.Fprintln(rep, user)
	}
}

func (c *IndexController) Index() {
	// c.Display()
	for key, value := range strings.Split(req.RequestURI, "/") {
		fmt.Fprintln(rep, key, value)
	}

}
