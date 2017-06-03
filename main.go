package main

import (
	"net/http"

	con "go.app.dt.com/controllers"
)

func main() {
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/image/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", con.IndexRouter)
	http.ListenAndServe(":6688", nil)
}
