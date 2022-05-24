package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (this *Person) PrintInfo() {
	fmt.Println("parent")
}

type Student struct {
	*Person
	id    int
	score float64
}

//func (this *Student) PrintInfo() {
//	fmt.Println("student")
//}

func main() {
	student := Student{&Person{"yorick", 'm', 20}, 1, 99}
	student.PrintInfo()
}
