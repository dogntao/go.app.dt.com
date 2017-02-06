package main

import (
	"fmt"
	"net/http"

	con "go.app.dt.com/app/controllers"
)

type DtHandler struct {
	port string
	con.BaseController
}

func main() {
	fmt.Println("start server")
	dtHandler := &DtHandler{port: ":8888"}
	http.ListenAndServe(dtHandler.port, dtHandler)
}
