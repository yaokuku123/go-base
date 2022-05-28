package main

import (
	"fmt"
	"net/rpc"
)

type LoginClient struct {
	c *rpc.Client
}

func (this *LoginClient) LoginService(name string, rsp *string) error {
	return this.c.Call("login.LoginService", name, rsp)
}

func InitClient(addr string) LoginClient {
	conn, _ := rpc.Dial("tcp", addr)
	return LoginClient{c: conn}
}

func main() {
	loginClient := InitClient("127.0.0.1:8080")
	var rsp string
	loginClient.LoginService("yorick_jun", &rsp)
	fmt.Println(rsp)
}
