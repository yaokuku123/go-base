package main

import "fmt"

func modify(slice []int) {
	slice[0] = 666
	fmt.Println("modify slice ", slice)
}

func main() {
	var slice1 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)

	var slice2 []int = make([]int, 3)
	fmt.Println(slice2)

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s[:])
	var s1 []int = s[2:5]
	fmt.Println(s1)
	var s2 []int = s1[2:6]
	fmt.Println(s2)

	var slice3 []int
	slice3 = append(slice3, 1)
	slice3 = append(slice3, 2, 3, 4)
	fmt.Println(slice3)

	slice4 := make([]int, 0, 1)
	c := cap(slice4)
	fmt.Println("c:", c)
	for i := 0; i < 50; i++ {
		slice4 = append(slice4, i)
		if n := cap(slice4); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}

	slice5 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	data1 := slice5[8:]
	data2 := slice5[:5]
	fmt.Println("data1:", data1)
	fmt.Println("data2:", data2)
	copy(data2, data1)
	fmt.Println(data2)
	fmt.Println(slice5)

	slice6 := []int{1, 2, 3, 4}
	modify(slice6)
	fmt.Println("main slice:", slice6)
}
