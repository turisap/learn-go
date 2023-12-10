package main

import "fmt"

func main() {

	var a = 5
	var b = 7.4

	fmt.Println(a, b)

	a = int(b)

	fmt.Println(a, b)

	var x string
	x = "ki"

	fmt.Println(x)

	var value int
	var price float64
	var done bool

	fmt.Println(value, price, done)

	var amount int

	fmt.Println(amount == 0.0)
}
