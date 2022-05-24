package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net dial err: ", err)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		fmt.Print("input data: ")
		fmt.Scan(&buf)
		conn.Write(buf)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err: ", err)
			return
		}
		fmt.Println("receive data: ", string(buf[:n]))
	}
}
