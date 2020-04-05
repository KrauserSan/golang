package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) foo4() {
	fmt.Println(p.first, p.last)
}
func main() {
	p1 := person{
		"krauser",
		"san",
		110,
	}
	p1.foo4()
}
