package main

import "fmt"

func main() {
	var numbers = [4]int{4, 5, -9, 100}

	fmt.Printf("%T\n", numbers)

	// SLICES
	var cities = []string{"london", "tokyo", "NY"}
	fmt.Printf("%T\n", cities)

	// MAP
	var balancies = map[string]float64{
		"USD": 30.4,
		"EUR": 16.5,
	}

	fmt.Printf("%T", balancies)

	type Person struct {
		name string
		age  int
	}

	var me Person

	fmt.Printf("%T\n", me)

	var x int = 2000000000000000

	fmt.Printf("%T\n", x)

	ptr := &x

	// prints out a memory address like this 0x1400009c030
	fmt.Printf("the pointer is of type %T and the value of %v\n", ptr, ptr)

	fmt.Printf("Type of f is %T\n", f)
}

func f() {

}
