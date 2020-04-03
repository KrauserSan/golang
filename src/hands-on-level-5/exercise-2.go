package main

import "fmt"

type people struct {
	first  string
	last   string
	flavor []string
}

func main() {
	p1 := people{
		first:  "krauser",
		last:   "san",
		flavor: []string{"choco", "vanilla"},
	}
	p2 := people{
		first:  "kintoki",
		last:   "chan",
		flavor: []string{"hazelnut", "coffee"},
	}
	mmap := map[string]people{}
	for _, v := range []people{p1, p2} {
		mmap[v.last] = v
	}
	fmt.Println(mmap)
}
