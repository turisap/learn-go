package main

import "fmt"

func main() {
	var slice1 []int
	var slice2 = []int{}

	// prints []int true
	fmt.Printf("%T %v\n", slice1, slice1 == nil)
	// prints []int false
	fmt.Printf("%T %v\n", slice1, slice2 == nil)

	var arr1 [3]int
	var arr2 = [3]int{}

	// prints [3]int
	fmt.Printf("%T \n", arr1)
	// prints [3]int
	fmt.Printf("%T \n", arr2)
}
