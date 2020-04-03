package main

import "fmt"

func main() {
	myslice := []int{1, 2, 3, 4, 5}
	for i, v := range myslice {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", myslice)
}
