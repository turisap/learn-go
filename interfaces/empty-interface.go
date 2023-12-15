package main

import "fmt"

// is analog of TS any
// used everywhere in definitions
type emptyInterface interface {
}

func main() {
	var a emptyInterface
	a = 5
	fmt.Println(a)

	a = "string"
	a = [3]int{1, 2, 3}
}
