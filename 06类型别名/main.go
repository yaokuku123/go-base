package main

import "fmt"

func main() {
	type bigint int64
	var x bigint = 100
	fmt.Println(x)
	fmt.Printf("type: %T\n", x)
}
