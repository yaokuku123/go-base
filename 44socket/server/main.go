package main

import (
	"fmt"
	"net"
	"strings"
)

func DealConn(conn net.Conn) {
	defer conn.Close()
	ipAddr := conn.RemoteAddr().String()
	fmt.Println(ipAddr, "connect success!")
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("connect read err: ", err)
			return
		}
		fmt.Println("read data: ", string(buf[:n]))
		if "exit" == string(buf[:n]) {
			fmt.Println(ipAddr, "logout")
			return
		}
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("server listen err: ", err)
		return
	}
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("server accept err: ", err)
			continue
		}
		go DealConn(conn)
	}
}
