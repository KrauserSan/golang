package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	sum := foo2(nums...)
	fmt.Println(sum)
	sumBar := bar2(nums)
	fmt.Println(sumBar)

}

func foo2(a ...int) int {
	tot := 0
	for i := 0; i < len(a); i++ {
		tot += a[i]
	}
	return tot
}

func bar2(a []int) int {
	tot := 0
	for _, v := range a {
		tot += v
	}
	return tot
}
