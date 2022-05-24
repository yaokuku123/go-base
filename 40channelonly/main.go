package main

import "fmt"

func Producer(channel chan<- int) {
	defer close(channel)
	for i := 0; i < 5; i++ {
		channel <- i
	}
}

func Consumer(channel <-chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println(<-channel)
	}
}

func main() {
	channel := make(chan int)
	go Producer(channel)
	Consumer(channel)
	fmt.Println("main finish")
}
