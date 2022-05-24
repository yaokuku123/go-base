package main

import "fmt"

func swap1(a, b int) {
	a, b = b, a
	fmt.Println("swap1 a=", a, " b=", b)
}

func swap2(a, b *int) {
	*a, *b = *b, *a
	fmt.Println("swap2 *a=", *a, " *b=", *b)
}

func main() {
	v := 10
	fmt.Printf("&v=%p\n", &v)
	var p *int = &v
	fmt.Printf("p=%p\n", p)
	*p = 100
	fmt.Println("*p=", *p, " v=", v)

	var p2 *int = new(int)
	*p2 = 110
	fmt.Println("*p2=", *p2)

	a, b := 10, 20
	swap1(a, b)
	fmt.Println("main a=", a, " b=", b)
	swap2(&a, &b)
	fmt.Println("main a=", a, " b=", b)
}
