package main

import "fmt"

func main() {
	s := struct {
		first string
		last  string
	}{
		first: "hello",
		last:  "world",
	}
	fmt.Println(s.first, s.last)

}
