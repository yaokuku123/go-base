package main

import "fmt"

func main() {
	var score int
	fmt.Print("input score: ")
	fmt.Scan(&score)
	switch score {
	case 90:
		fmt.Println("very good")
	case 80:
		fmt.Println("good")
	case 70:
		fmt.Println("not good")
	default:
		fmt.Println("bad")
	}

	switch {
	case score >= 90:
		fmt.Println("very good 2")
	case score >= 80:
		fmt.Println("good 2")
	default:
		fmt.Println("bad 2")
	}
}
