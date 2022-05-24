package main

import "fmt"

func main() {
	if a := 3; a == 3 {
		fmt.Println("a = 3")
	}

	var score int
	fmt.Print("input score: ")
	fmt.Scan(&score)
	if score > 80 {
		fmt.Println("very good")
	} else if score > 60 {
		fmt.Println("good")
	} else {
		fmt.Println("not good")
	}
}
