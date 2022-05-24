package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("input num1: ")
	fmt.Scanf("%d", &num1)
	fmt.Print("input num2: ")
	fmt.Scan(&num2)
	fmt.Println(num1 + num2)
}
