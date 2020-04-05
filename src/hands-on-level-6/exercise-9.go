package main

import (
	"fmt"
)

func foo9() string {
	return fmt.Sprint("func from foo8")
}

func bar9(f func() string, s string) {
	fmt.Print(f())
	fmt.Print("\t", s)
}

func main() {
	bar9(foo9, "inside bar9")
}
