package main

import "fmt"

func main() {
	fmt.Println(`He says "Hello go!"`)
	fmt.Println("he says \\go")

	str := "Hello go"
	// prints out 72 as ASCI code
	// because strings are slices
	// of byte arrays
	fmt.Println(str[0])

	// error because strings are
	// immutable
	//str[0] = 34

	// prints out
	// He says \n "hello" (raw string)
	// no newline
	fmt.Println(`He says \n "hello"`)
}
