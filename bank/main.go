package main

import (
	"fmt"
	"mypackages/numbers"
	"strings"
)

var n [10]float64

func init() {
	// use init() to
	// initialize package-scoped vars
	for i, _ := range n {
		n[i] = float64(0.8) * 0.8
	}
}

func main() {
	fmt.Println(numbers.Odd(2))
	fmt.Println(numbers.Kiskisk(2))
	sayHello()

	f := toFarenheit(boilingPoint)
	fmt.Println("Water boiling point in F is:", f)

	fmt.Println("\n" + strings.Repeat("@", 20))

	fmt.Println(n)
}
