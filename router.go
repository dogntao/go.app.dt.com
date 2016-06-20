package main

import (
	"fmt"
	"net/http"
)

func RouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "thi is router")
}
