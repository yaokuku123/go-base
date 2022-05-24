package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)
	timeout := make(chan bool)
	go func() {
		for {
			select {
			case value := <-channel:
				fmt.Println(value)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				timeout <- true
				break
			}
		}
	}()
	time.Sleep(time.Second)
	channel <- 233
	<-timeout
}
