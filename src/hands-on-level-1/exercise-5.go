package main

import "fmt"

type mytype int
var x mytype
var y int


func main(){
  fmt.Printf("%T\t%v", x,x)
  x = 42
  fmt.Printf("%T\t%v", x,x)
  y = int(x)
  fmt.Printf("%T\t%v", y, y)
}
