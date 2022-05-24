package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person
	id    int
	score float64
}

type Superman struct {
	*Person
	id    int
	power string
}

func main() {
	student1 := Student{Person{"yorick", 'm', 24}, 1, 95.5}
	fmt.Println(student1)

	var student2 Student
	student2.name = "tom"
	student2.sex = 'm'
	student2.age = 20
	student2.id = 2
	student2.score = 79.9
	fmt.Println(student2)

	var student3 Student
	student3.Person = Person{"jerry", 'm', 18}
	student3.id = 3
	student3.score = 60.0
	fmt.Println(student3)

	superman := Superman{&Person{"yorick", 'm', 24}, 110, "super power"}
	fmt.Println(superman)
	fmt.Printf("%+v\n", superman)
	fmt.Println(superman.name, superman.age, superman.sex)
}
