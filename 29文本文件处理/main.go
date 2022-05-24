package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "hello world go go go!"
	fmt.Println(strings.Contains(str1, "hello"))
	fmt.Println(strings.Contains(str1, "nihao"))

	strList1 := []string{"tom", "jerry", "yorick"}
	fmt.Println(strings.Join(strList1, ","))

	fmt.Println(strings.Index(str1, "go"))
	fmt.Println(strings.Index(str1, "og"))

	fmt.Println(strings.Repeat("yo", 3))

	fmt.Println(strings.Replace(str1, "go", "java", 2))
	fmt.Println(strings.Replace(str1, "go", "java", -1))

	fmt.Println(strings.Split(str1, " "))

	str2 := "   hello go   "
	fmt.Println(len(str2))
	trimStr2 := strings.Trim(str2, " ")
	fmt.Println(trimStr2)
	fmt.Println(len(trimStr2))
}
