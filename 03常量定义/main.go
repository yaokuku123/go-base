package main

import "fmt"

const (
	x = iota
	y
	z
	w
)
const v = iota
const (
	h, i, j = iota, iota, iota
	k, l    = iota, iota
)

func main() {
	const PI float64 = 3.14

	fmt.Println(x, y, z, w, PI, v, h, i, j, k, l)
}
