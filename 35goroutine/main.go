package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Println("newTask i = ", i)
		time.Sleep(time.Second)
	}
}

func main() {

	go newTask()

	i := 0
	for {
		i++
		fmt.Println("main i = ", i)
		time.Sleep(time.Second)
		if i == 5 {
			break
		}
	}
}
