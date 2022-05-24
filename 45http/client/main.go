package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8080/go")
	if err != nil {
		fmt.Println("get err: ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("header = ", resp.Header)
	fmt.Printf("resp status %s\nstatusCode %d\n", resp.Status, resp.StatusCode)
	fmt.Printf("body type = %T\n", resp.Body)
	buf := make([]byte, 1024)
	var tmp string
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}
		if n == 0 {
			fmt.Println("读取内容结束")
			break
		}
		tmp += string(buf[:n]) //累加读取的内容
	}
	fmt.Println("tmp=", tmp)
}
