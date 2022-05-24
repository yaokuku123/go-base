package main

import (
	"fmt"
	"regexp"
)

func main() {
	context := "3.14 123123 .68 haha 1.0 abc 6.66 123."
	expr1 := regexp.MustCompile(`\d+\.\d+`)
	res1 := expr1.FindAllStringSubmatch(context, -1)
	fmt.Println(res1)

}
