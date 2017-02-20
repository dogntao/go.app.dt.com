package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// type DtHandler struct {
// 	port string
// 	con.Router
// }

// func main() {
// 	fmt.Println("start server")
// 	dtHandler := &DtHandler{port: ":8888"}
// 	http.ListenAndServe(dtHandler.port, dtHandler)
// }

type User struct {
	Name  string
	Hobby [3]string
}

// 路由转发(/Index/index)
func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Printf("%s,%s \r\n", k, v[0])
	}
	Users := make([]User, 3)
	Users[0].Name = "dongtao"
	Users[0].Hobby[0] = "swimming"
	Users[0].Hobby[1] = "basketball"
	Users[1].Name = "dt"
	Users[1].Hobby[0] = "football"
	Users[1].Hobby[1] = "pingpang"

	t, err := template.ParseFiles("view/layouts/header.html", "view/index/index.html", "view/layouts/footer.html")
	t.ExecuteTemplate(w, "index", Users)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8888", nil)
}
