package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	p.first = "something"
}

func main() {
	p1 := person{
		"krauser",
		"san",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
