package main

import "fmt"

func main() {
	var ch byte = 97
	var a int = int(ch)
	fmt.Println(a)

	var str string = "你好"
	var arrStr []byte = []byte(str)
	fmt.Println(arrStr)
}
