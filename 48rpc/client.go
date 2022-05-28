package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:8080")
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
