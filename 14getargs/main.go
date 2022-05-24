package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if args == nil || len(args) < 3 {
		fmt.Println("err: xxx ip port")
		return
	}
	ip := args[1]
	port := args[2]
	fmt.Println("ip=", ip, " port=", port)
}
