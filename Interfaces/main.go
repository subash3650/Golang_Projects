package main

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
}
type rectange struct {
	width  float64
	height float64
}
type circle struct {
	radius float64
}

func (r rectange) Area() float64 {
	return r.width * r.height
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func printArea(s shape) {
	fmt.Printf("The area is %.2f\n", s.Area())
}

func main() {
	r := rectange{width: 5, height: 3}
	c := circle{radius: 4}

	printArea(r)
	printArea(c)
}
