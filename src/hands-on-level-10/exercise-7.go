// package main
//
// import (
// 	"fmt"
// 	"runtime"
// )
//
// func main() {
// 	c := make(chan int)
// 	for i := 0; i < 10; i++ {
// 		go func(ind int) {
// 			for j := 0; j < 10; j++ {
// 				c <- ind*10 + j
// 			}
// 			if runtime.NumGoroutine() == 2 {
// 				close(c)
// 			}
//
// 		}(i)
// 	}
// 	for v := range c {
// 		fmt.Println(v)
// 	}
//
// 	fmt.Println("Exiting program")
// }

package main

import "fmt"

func main() {
	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("about to exit")
}
