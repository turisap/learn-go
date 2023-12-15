package main

import "fmt"

func main() {
	// IIFE analog
	func(msg string) {
		fmt.Println(msg)
	}("I am anonymous fn")

	a := increment(10)

	// prints out "func() int"
	fmt.Printf("%T\n", a)

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}
