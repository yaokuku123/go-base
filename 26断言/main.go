package main

import "fmt"

type Human interface {
	SayHello()
}

type Person struct {
	name string
	age  int
}

func (this *Person) SayHello() {
	fmt.Println("person say hello")
}

func main() {
	//TODO interface接口的万能类型视频
	list := make([]interface{}, 3)
	list[0] = 10
	list[1] = "Hello"
	list[2] = Person{"yorick", 24}
	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Println("int index=", index, " value=", value)
		} else if value, ok := element.(string); ok {
			fmt.Println("string index=", index, " value=", value)
		} else if value, ok := element.(Person); ok {
			fmt.Println("Person index=", index, " value=", value)
		}
	}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Println("int index=", index, " value=", value)
		case string:
			fmt.Println("string index=", index, " value=", value)
		case Person:
			fmt.Println("Person index=", index, " value=", value)
		}
	}

	var num interface{} = 10
	var str interface{} = "hello world"
	var human interface{} = &Person{name: "yorick", age: 25}
	_, ok1 := num.(int)
	if ok1 {
		fmt.Println("num is int type")
	}
	_, ok2 := str.(string)
	if ok2 {
		fmt.Println("str is string type")
	}
	_, ok3 := human.(*Person)
	if ok3 {
		fmt.Println("human is *Person type")
	}
	_, ok4 := human.(Human)
	if ok4 {
		fmt.Println("human is Human type")
	}
}
