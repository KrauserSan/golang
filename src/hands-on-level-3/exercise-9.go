package main

import "fmt"

func main() {
	favSport := "sport"
	switch favSport {
	case "sPort":
		fmt.Println("this doesn't works")
	default:
		fmt.Println("This is default")
	}

}
