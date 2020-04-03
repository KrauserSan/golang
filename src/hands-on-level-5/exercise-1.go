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
	for i, v := range []people{p1, p2} {
		fmt.Println(i, v)
		for ind, val := range v.flavor {
			fmt.Printf("%d\t%v\n", ind, val)
		}
	}
}
