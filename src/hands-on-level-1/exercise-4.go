package main

import "fmt"

type mytype int
var x mytype

func main(){
  fmt.Printf("%T\t%v", x,x)
  x = 42
  fmt.Printf("%T\t%v", x,x)
}
