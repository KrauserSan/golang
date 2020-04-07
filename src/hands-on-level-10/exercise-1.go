package main

import (
	"fmt"
)

func main() {
	// USING FUNC LITERAL
	// c := make(chan int)
	//
	// // By default sends and receives block until both the sender and receiver are ready.
	// // sender
	// go func(ch chan<- int) {
	// 	ch <- 42
	// }(c)
	// // receiver
	// fmt.Println(<-c)

	// USING BUFFERED CHAN
	c := make(chan int, 1)
	c <- 41
	fmt.Println(<-c)

}
