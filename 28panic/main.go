package main

import (
	"fmt"
)

func TestA() {
	fmt.Println("func TestA")
}

func TestB() (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("internal error: %v", x)
		}
	}()
	panic("func TestB panic")
}

func TestC() {
	fmt.Println("func TestC")
}

func main() {
	//TODO 视频查看这块panic
	TestA()
	err := TestB()
	fmt.Println(err)
	TestC()
}
