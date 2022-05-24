package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	str := "yorick"
	for index, value := range str {
		fmt.Printf("index=%d value=%c\n", index, value)
	}
}
