package main

import "fmt"

func main() {
	name := "Kirill"
	fmt.Println("kis", name)

	a, b := 4, 6

	fmt.Println("Sum:", a+b)

	fmt.Printf("Your age is %d\n", 33)
	fmt.Printf("X is %d, Y is %f", 5, 8.5)

	figure := "Circle"
	radius := 5
	pi := 3.14

	_, _ = pi, figure

	fmt.Println('\n')
	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Radius is %+d\n", radius)
	fmt.Printf("PI is: %+f\n", pi)
	fmt.Printf("The diameter of %s with radius of %d is %id", figure, radius, 2*radius)

	fmt.Printf("This can print %v\n", figure)

	text := "new"
	fmt.Printf("This can print %T\n", text)
	// This can print string

	closed := true
	fmt.Printf("File is closed %t\n", closed)
}
