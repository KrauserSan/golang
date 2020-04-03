package main

import "fmt"

func main() {
	for x := 65; x <= 90; x++ {
		for y := 0; y < 3; y++ {
			fmt.Printf("%#U\n", x)
		}
		fmt.Println(string(x))

	}

}
