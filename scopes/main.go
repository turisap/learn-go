package main

import (
	"fmt"
)

const done = true // package-scoped

// package-scoped does not throw the unused var error
// as it can be used in other files
var count = 10

func main() {
	const done = false // block-scoped
	const i = 4
	const a float64 = i
	const b int8 = i

	// shadowing happens
	fmt.Printf("done local scoped is: %v \n", done)

	f1()
}

func f1() {
	fmt.Printf("done package scoped is: %v \n", done)
}
