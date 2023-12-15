package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "kirill"
	// prints a memory address
	// like 0x1400011e010
	fmt.Println(&name)

	var x int = 4
	p := &x

	// prints out *int, 0x1400009c018
	fmt.Printf("%T, %v\n", p, p)
	fmt.Printf("Address of x is %p\n", p)

	var y int = 4
	p = &y
	// address of the pointer itself
	fmt.Println(&p)
	fmt.Println(p)

	fmt.Println("\n" + strings.Repeat("@", 20))

	x = 100 //int
	// pointer type definition
	// depends on the var value
	var pointer *int = &x

	//*pointer = *pointer / 2
	// prints out 50
	fmt.Println(*pointer)

	// same as x = 33
	*pointer = 33
	// prints 33, 33
	fmt.Println(x, *pointer)
	// prints true
	fmt.Println(x == *pointer)
}
