package main

import (
	"fmt"
)

func foo8() func() {
	return func() {
		fmt.Println("func from foo8")
	}
}

func main() {
	x := foo8()
	x()
}
