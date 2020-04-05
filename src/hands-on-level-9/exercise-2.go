package main

import (
	"fmt"
)

type person struct {
	First, Last string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(p.First, (p).Last)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		"krauser", "san",
	}

	saySomething(&p1)
	p1.speak()
	// saySomething(p1)
	// cannot use p1 (type person) as type human in argument to saySomething:
	// 	person does not implement human (speak method has pointer receiver)
}
