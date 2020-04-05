package main

import (
	"fmt"
)

func main() {
	x := 53
	fmt.Printf("Type - %T\tValue - %v", x, x)
	fmt.Printf("Type - %T\tValue - %v", &x, &x)
}
