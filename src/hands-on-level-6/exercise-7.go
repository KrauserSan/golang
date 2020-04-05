package main

import (
	"fmt"
)

func main() {
	a := func() {
		fmt.Println("hello, world!")
	}
	fmt.Printf("%T\n", a)
	a()
}
