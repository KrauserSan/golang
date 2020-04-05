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
	var lock sync.Mutex
	wg.Add(goCount)
	for i := 0; i < goCount; i++ {
		go func() {
			lock.Lock()
			tmp := count
			runtime.Gosched()
			tmp++
			count = tmp
			fmt.Println(count)
			lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final value", count)
}
