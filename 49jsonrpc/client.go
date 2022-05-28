package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

// 使用nc模拟server：nc -l 127.0.0.1 8080
// 发送的rpc json格式：{"method":"login.LoginService","params":["yorick"],"id":0}
func main() {
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("rpc Dial err:", err)
		return
	}
	defer conn.Close()
	var rsp string
	err = conn.Call("login.LoginService", "yorick", &rsp)
	if err != nil {
		fmt.Println("login.LoginService Call err:", err)
		return
	}
	fmt.Println(rsp)
}
