package main

import (
	"fmt"
	"net/http"
	//	"net/url"

	con "go.app.dt.com/controller"
	//	"strings"
)

type DtHandler struct {
	port string
}

func (dt *DtHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	con.Base(w, r)
	return
}

func main() {
	fmt.Println("start server")
	dtHandler := &DtHandler{port: ":8888"}
	http.ListenAndServe(dtHandler.port, dtHandler)
}
