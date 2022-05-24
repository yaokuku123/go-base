package main

import "fmt"

func main() {
	i := 10
	str := "yorick"

	// method 1
	f1 := func() {
		fmt.Println("method1 i=", i, " str=", str)
	}
	f1()

	// method 2
	func() {
		fmt.Println("method2 i=", i, " str=", str)
	}()

	// method 3
	res := func(num1, num2 int) (res int) {
		res = num1 + num2
		return
	}(10, 20)
	fmt.Println("method3 res=", res)

	// method4
	func(num1, num2 int) (res int) {
		res = num1 + num2
		fmt.Println("method4 res=", res)
		return
	}(30, 40)
}
