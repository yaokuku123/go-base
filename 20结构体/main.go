package main

import (
	"fmt"
	"github.com/yaokuku123/go-base/20结构体/entity"
)

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func modify1(student Student) {
	student.name = "yorick_jun"
	fmt.Println("modify1 student: ", student)
}

func modify2(student *Student) {
	student.name = "yorick_jun2"
	fmt.Println("modify2 student: ", *student)
}

func main() {
	var student1 Student = Student{1, "yorick", 'm', 18, "bj"}
	fmt.Println(student1)

	var student2 *Student = &Student{2, "tom", 'w', 20, "sh"}
	fmt.Println(*student2)

	var student3 Student
	student3.id = 3
	student3.name = "jerry"
	student3.sex = 'm'
	student3.age = 11
	student3.addr = "tj"
	fmt.Println(student3)

	var p3 *Student = &student3
	p3.id = 4
	p3.name = "jerry2"
	fmt.Println(*p3, student3)

	student4 := Student{4, "smith", 'm', 21, "yn"}
	modify1(student4)
	fmt.Println("main student: ", student4)
	modify2(&student4)
	fmt.Println("main student: ", student4)

	teacher1 := entity.Teacher{10, "ZHU"}
	fmt.Println(teacher1)
}
