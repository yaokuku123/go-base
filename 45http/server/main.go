package main

import (
	"fmt"
	"net/http"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println(r.RemoteAddr, " connect success")
	fmt.Println("method=", r.Method)
	fmt.Println("url=", r.URL.Path)
	fmt.Println("header=", r.Header)
	fmt.Println("body=", r.Body)
	fmt.Fprint(w, "hello world")
}

func main() {
	http.HandleFunc("/go", MyHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
