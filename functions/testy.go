package main

import "fmt"

func main() {
	var s []int

	// non-init value of a slice
	// prints out []int(nil)
	fmt.Printf("%#v", s)
}
