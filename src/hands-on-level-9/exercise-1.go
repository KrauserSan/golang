package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(a int) {
			fmt.Println("anonymous func", a)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
