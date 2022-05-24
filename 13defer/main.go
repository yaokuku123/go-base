package main

import "fmt"

func Div(x int) {
	fmt.Println(1 / x)
}

func main() {
	//this is a defer2
	//this is a defer1
	//defer fmt.Println("this is a defer1")
	//defer fmt.Println("this is a defer2")
	//Div(0)
	//defer fmt.Println("this is a defer3")
	//fmt.Println("this is a main")

	//this is a main
	//this is a defer3
	//this is a defer2
	//this is a defer1
	//defer fmt.Println("this is a defer1")
	//defer fmt.Println("this is a defer2")
	//defer Div(0)
	//defer fmt.Println("this is a defer3")
	//fmt.Println("this is a main")

	a, b := 10, 20
	defer func(a int) {
		fmt.Println("defer a=", a, " b=", b)
	}(a)
	a += 10
	b += 10
	fmt.Println("main a=", a, " b=", b)
}
