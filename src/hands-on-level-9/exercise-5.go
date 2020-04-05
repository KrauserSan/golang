package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	const goCount = 100
	var count int64
	var wg sync.WaitGroup
	wg.Add(goCount)
	for i := 0; i < goCount; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&count))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final value", count)
}
