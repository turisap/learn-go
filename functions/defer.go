package main

import "fmt"

func foo() {
	fmt.Println("this is foo")
}
func bar() {
	fmt.Println("this is bar")
}
func foobar() {
	fmt.Println("this is fooBar")
}

func main() {
	// 4
	defer foo()
	// 1
	bar()

	// 2
	fmt.Println("Just a string after defer foo and calling bar")

	// 3
	defer foobar()
}
