package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 3)
	go func() {
		defer fmt.Println("child finish")
		for i := 0; i < 3; i++ {
			fmt.Println("child running,i=", i)
			channel <- i
		}
	}()
	time.Sleep(time.Second * 2)
	for i := 0; i < 3; i++ {
		res := <-channel
		fmt.Println("main running,res=", res)
	}
	fmt.Println("main finish")
}
