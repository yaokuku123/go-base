package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 提供rpc服务的对象
type LoginController struct {
}

// 提供rpc的函数
// 1 必须可被外界调用 Public
// 2 必须有两个参数，分别为传入参数和传出参数
// 3 传出参数必须为指针类型
// 4 只能有一个返回值为error
func (this *LoginController) LoginService(name string, rsp *string) error {
	*rsp = name + " login success!"
	return nil
}

// 使用nc模拟client: echo -e '{"method":"login.LoginService","params":["yorick"],"id":0}' | nc localhost 8080
// 返回的json结果：{"id":0,"result":"yorick login success!","error":null}
func main() {
	// 1 注册rpc对象
	err := rpc.RegisterName("login", new(LoginController))
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
		jsonrpc.ServeConn(conn)
	}
}
