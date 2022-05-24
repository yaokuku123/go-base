package main

import "fmt"

type Humaner interface {
	SayHi()
}
type Student struct {
	name string
}
type Teacher struct {
	name string
}

func (this *Student) SayHi() {
	fmt.Println("student", this.name, " say hi")
}

func (this *Teacher) SayHi() {
	fmt.Println("teacher", this.name, " say hi")
}

func WhoSayHi(humaner Humaner) {
	humaner.SayHi()
}

func main() {
	var humaner Humaner
	var s Student = Student{"yorick"}
	var t *Teacher = &Teacher{"ZHU"}
	humaner = &s
	humaner.SayHi()
	humaner = t
	humaner.SayHi()
	WhoSayHi(&s)
	WhoSayHi(t)
}
