package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	const goCount = 100
	var count int
	var wg sync.WaitGroup
	wg.Add(goCount)
	for i := 0; i < goCount; i++ {
		go func() {
			tmp := count
			runtime.Gosched()
			tmp++
			count = tmp
			fmt.Println(count)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final value", count)
}
