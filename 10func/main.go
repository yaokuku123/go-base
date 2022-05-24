package main

import (
	"fmt"
)

func test01() {
	fmt.Println("test01")
	fmt.Println("================")
}

func test02(v1 int, v2 string) {
	fmt.Println("test02 v1=", v1, " v2=", v2)
	fmt.Println("================")
}

func test03(args ...int) {
	for index, value := range args {
		fmt.Println("index=", index, " value=", value)
	}
	fmt.Println("================")
}

func test04() (num1 int, num2 int) {
	num1 = 20
	num2 = 10
	return
}

func test05(num1, num2 int) (min, max int) {
	if num1 > num2 {
		min = num2
		max = num1
	} else {
		min = num1
		max = num2
	}
	return
}

func test06(num int) int {
	if num == 100 {
		return num
	}
	return num + test06(num+1)
}

type FuncType func(int, int) int

func Calc(num1, num2 int, f FuncType) (res int) {
	res = f(num1, num2)
	return
}
func Add(num1, num2 int) (res int) {
	res = num1 + num2
	return
}
func Sub(num1, num2 int) (res int) {
	res = num1 - num2
	return
}

func main() {
	test01()

	test02(100, "yorick")

	test03(10, 20, 30, 40)

	num1, num2 := test04()
	fmt.Println("sum= ", num1+num2)
	fmt.Println("================")

	min, max := test05(num1, num2)
	fmt.Println("min=", min, " max=", max)
	fmt.Println("================")

	sum := test06(1)
	fmt.Println("sum=", sum)
	fmt.Println("================")

	res1 := Calc(10, 20, Add)
	res2 := Calc(10, 20, Sub)
	fmt.Println("res1=", res1, " res2=", res2)
	fmt.Println("================")
}
