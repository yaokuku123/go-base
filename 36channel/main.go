package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go func() {
		defer fmt.Println("child process finish")
		fmt.Println("child process running")
		time.Sleep(time.Second * 2)
		channel <- "ok"
	}()

	res := <-channel
	fmt.Println(res)
	fmt.Println("main process finish")
}
