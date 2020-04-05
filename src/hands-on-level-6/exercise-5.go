package main

import (
	"fmt"
	"math"
)

type square struct {
	length  float64
	breadth float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return 3.14 * (math.Pow(c.radius, 2))
}

func (s square) area() float64 {
	return s.length * s.breadth
}

func info(sh shape) {
	fmt.Println(sh.area())
}
func main() {
	square1 := square{
		5,
		6,
	}
	circle1 := circle{
		4,
	}
	info(square1)
	info(circle1)
}
