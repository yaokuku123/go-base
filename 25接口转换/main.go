package main

import "fmt"

type Humaner interface {
	Say()
}
type Personer interface {
	Humaner
	Sing()
}
type Student struct {
	name string
}

func (this *Student) Say() {
	fmt.Println("student say")
}

func (this *Student) Sing() {
	fmt.Println("student sing")
}

func main() {
	var humaner Humaner
	var personer Personer
	personer = &Student{"yorick"}
	personer.Say()
	personer.Sing()
	humaner = personer
	humaner.Say()

}
