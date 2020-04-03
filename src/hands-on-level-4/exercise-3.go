package main

import "fmt"

func main() {
	myslice := []int{1, 2, 3, 4, 5}
	for i, v := range myslice {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", myslice)

	ms2 := myslice[1:4] //  only assigns a pointer and changes the len&cap. Modifying ms2 will update myslice as well
	fmt.Println(ms2)
}
