package main

import (
	"fmt"
	"net/http"
	//	"net/url"

	con "go.app.dt.com/controller"
	//	"strings"
)

func main() {
	fmt.Println("start server")
	http.HandleFunc("/", con.Base)
	http.ListenAndServe(":8888", nil)
}
