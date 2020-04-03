package main

import "fmt"

func main() {
	ms := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	fmt.Println(ms)
	for row := 0; row < len(ms); row++ {
		for col := 0; col < len(ms[0]); col++ {
			fmt.Println(ms[row][col])
		}
	}
}
