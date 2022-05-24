package main

import "fmt"

func modify1(arr [5]int) {
	arr[0] = 233
	fmt.Println("modify1 arr=", arr)
}

func modify2(arr *[5]int) {
	arr[0] = 666
	fmt.Println("modify2 arr=", *arr)
}

func main() {
	var arr [10]int
	fmt.Println(len(arr), cap(arr))
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	for i, v := range arr {
		fmt.Println("i=", i, " v=", v)
	}

	arr2 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Print(arr2[i][j], " ")
		}
		fmt.Println()
	}

	var arr3 [5]int = [5]int{1, 2, 3, 4, 5}
	modify1(arr3)
	fmt.Println("main arr=", arr3)

	modify2(&arr3)
	fmt.Println("main arr=", arr3)
}
