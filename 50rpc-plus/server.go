package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 指定实现的rpc接口
type LoginInterface interface {
	LoginService(string, *string) error
}

// 提供rpc服务的对象
type LoginController struct {
}

func (this *LoginController) LoginService(name string, rsp *string) error {
	*rsp = name + " login success!"
	return nil
}

func RegisterLoginService(i LoginInterface) error {
	return rpc.RegisterName("login", i)
}

func main() {
	// 1 注册rpc对象
	err := RegisterLoginService(new(LoginController))
	if err != nil {
		fmt.Println("rpc RegisterName err:", err)
		return
	}
	// 2 设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net Listen err:", err)
		return
	}
	defer listener.Close()
	for {
		// 3 建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener Accept err:", err)
			return
		}
		// 4 绑定服务
		rpc.ServeConn(conn)
	}
}
