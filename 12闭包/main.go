package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	i := 100
	str := "yorick"

	func() {
		i = 200
		str = "yorick233"
		fmt.Println("func i=", i, " str=", str)
	}()
	fmt.Println("main i=", i, " str=", str)

	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
