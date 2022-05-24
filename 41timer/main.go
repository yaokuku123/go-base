package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.NewTimer(time.Second * 2)
	t1 := time.Now()
	fmt.Println(t1)
	t2 := <-time1.C
	fmt.Println(t2)

	time.Sleep(time.Second * 2)
	fmt.Println("after 2s")

	<-time.After(time.Second * 2)
	fmt.Println("after 2s too")
}
