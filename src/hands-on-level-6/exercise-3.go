package main

import (
	"fmt"
)

func main() {
	defer foo3()
	fmt.Println("inside main")

}

func foo3() {
	fmt.Println("inside foo3")
}
