package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sq(4.3))

	// @COOL captured variables from returning tuple
	s, prod := sumAndProd(4)

	fmt.Println(s, prod)

	fmt.Println("sum 3", sum(3))
}

func sq(a float64) float64 {
	return math.Pow(a, 2.0)
}

func sumAndProd(a int) (int, int) {
	return a + a, a * a
}

// named return value
func sum(a int) (s int) {
	// prints out 0, zero-init
	fmt.Println("S is", s)

	s = a + a

	// naked return, returning s
	return
}

func trt(a, b int) int {
	return a + b
}
