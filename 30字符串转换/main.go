package main

import (
	"fmt"
	"strconv"
)

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	byteArr := make([]byte, 0, 10)
	byteArr = strconv.AppendInt(byteArr, 233, 10)
	byteArr = strconv.AppendBool(byteArr, true)
	byteArr = strconv.AppendQuote(byteArr, "hello yorick")
	fmt.Println(string(byteArr))

	a := strconv.FormatBool(true)
	b := strconv.FormatInt(666, 10)
	c := strconv.Itoa(21)
	fmt.Println(a, b, c)
	fmt.Printf("a:%T,b:%T,c:%T\n", a, b, c)

	a1, err := strconv.ParseBool("true")
	checkError(err)
	b1, err := strconv.ParseInt("12345", 10, 64)
	checkError(err)
	c1, err := strconv.ParseFloat("12.12", 64)
	checkError(err)
	fmt.Println(a1, b1, c1)

}
