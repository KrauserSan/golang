package main

import (
	"fmt"
)

func main() {
	f := foo()
	b1, b2 := bar()
	fmt.Println(f, b1, b2)

}

func foo() int {
	return 2
}

func bar() (int, string) {
	return 4, "hello, world!"
}
