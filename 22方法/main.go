package main

import "fmt"

type MyInt int

func (a MyInt) Add(b MyInt) MyInt {
	return a + b
}

func Add(a, b MyInt) MyInt {
	return a + b
}

type Person struct {
	name string
	sex  byte
	age  int
}

func (this Person) PrintInfo() {
	fmt.Println(this.name, this.sex, this.age)
}

func (this Person) SetInfoValue() {
	this.name = "yorick_jun"
	this.sex = 'f'
	this.age = 100
	fmt.Println("SetInfoValue: ", this)
}

func (this *Person) SetInfoPointer() {
	this.name = "yorick_jun_pointer"
	this.sex = 'm'
	this.age = 101
	fmt.Println("SetInfoPointer: ", *this)
}

func main() {
	sum1 := Add(10, 20)
	fmt.Println(sum1)
	var a MyInt
	a = 30
	sum2 := a.Add(40)
	fmt.Println(sum2)

	person1 := Person{"yorick", 'm', 23}
	person1.PrintInfo()

	person2 := Person{"yoyo", 'f', 22}
	person2.SetInfoValue()
	fmt.Println("main1: ", person2)
	person2.SetInfoPointer()
	fmt.Println("main2: ", person2)

	//TODO 这块看下视频，指针的对象调用方法
	var person3 *Person = &Person{"tom", 'm', 18}
	person3.SetInfoValue()
	fmt.Println("main3: ", *person3)
	person3.SetInfoPointer()
	fmt.Println("main4: ", *person3)

}
