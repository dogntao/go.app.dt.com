package main

import (
	"fmt"
	"net/http"
	"net/url"
	//	"strings"
)

func base(w http.ResponseWriter, r *http.Request) {
	urlAdd := "https://www.baidu.com/s?wd=%E6%90%9C%E7%B4%A2&rsv_spt=1&issp=1&f=8&rsv_bp=0&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_sug3=7&rsv_sug1=6"
	urlInfo, err := url.ParseRequestURI(urlAdd)
	if err != nil {
		panic(err)
	}
	fmt.Println(urlInfo.Query())
}

func main() {
	fmt.Println("start server")
	mux := http.NewServeMux()
	mux.HandleFunc("/", base)
	http.ListenAndServe(":8888", mux)
}
