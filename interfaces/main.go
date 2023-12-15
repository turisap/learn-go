package main

import (
	"fmt"
	"math"
	"strings"
)

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * math.Pi * 2
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * 3
}

func (c rectangle) area() float64 {
	return c.width * c.height
}

func info(s shape) {
	fmt.Println(s.area())
}

type shape interface {
	area() float64
}

func main() {
	var s shape
	// prints out nil
	fmt.Println(s)

	c := circle{radius: 12}
	s = c
	// assigning a type to an interface
	// prints {12}
	fmt.Println(s)

	r := rectangle{width: 4, height: 3}

	// only accepts shapes
	info(c)
	info(r)

	fmt.Println("\n" + strings.Repeat("@", 20))
	assertions()
}

/* TYPE ASSERTIONS */
func assertions() {
	var s shape = circle{radius: 12}

	s = rectangle{width: 1, height: 3}

	ball, ok := s.(circle)

	if ok == true {
		fmt.Println(ball, ok)
	}

	switch s.(type) {
	case circle:
		fmt.Println("shape is a circle")

	case rectangle:
		fmt.Println("shape is a rectangle")
	}
}
