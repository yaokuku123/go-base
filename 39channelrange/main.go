package main

import "fmt"

func main() {
	channel := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			channel <- i
		}
		close(channel)
	}()
	for data := range channel {
		fmt.Println(data)
	}
	fmt.Println("finish")
}
