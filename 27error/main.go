package main

import (
	"errors"
	"fmt"
)

func Divide(a, b float64) (res float64, err error) {
	if b == 0 {
		res = 0
		err = errors.New("runtime error: divide by zero")
		return
	}
	res = a / b
	err = nil
	return
}

func main() {
	var err1 error = errors.New("a normal err1")
	fmt.Println(err1)

	var err2 error = fmt.Errorf("%s", "a normal err2")
	fmt.Println(err2)

	r, err := Divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}
