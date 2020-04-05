package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOARCH, runtime.GOOS, runtime.NumCPU(), runtime.NumGoroutine())
}
