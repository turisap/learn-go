package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return c.radius * math.Pi * 2
}

func (c rectangle) area() float64 {
	return c.width * c.height
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c := circle{radius: 12}
	r := rectangle{width: 4, height: 3}

	info(c)
	info(r)
}
