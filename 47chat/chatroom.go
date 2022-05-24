package main

import (
	"fmt"
	"net"
)

type User struct {
	id   string
	name string
	msg  chan string
}

var allUsers = make(map[string]User)

var message = make(chan string, 10)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	go broadcast()
	fmt.Println("服务器启动成功...")

	for {
		fmt.Println("监听中...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	for {
		fmt.Println("具体处理...")
		remoteAddr := conn.RemoteAddr().String()
		user := User{
			id:   remoteAddr,
			name: remoteAddr,
			msg:  make(chan string),
		}
		allUsers[user.id] = user
		buf := make([]byte, 1024)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		// cnt - 1 去掉最后的回车换行
		fmt.Println("获取客户端数据：", string(buf[:cnt-1]), " cnt:", cnt)
	}
}

func broadcast() {
	fmt.Println("启动广播go程成功...")
	info := <-message
	for _, user := range allUsers {
		user.msg <- info
	}
}
