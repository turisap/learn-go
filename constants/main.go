package main

import "fmt"

func main() {

	const (
		min1 = 300
		min2 = 500
		min3
	)

	// prints 300, 500, 500 because the last const takes
	// its value from the previous one
	fmt.Println("Constants are:", min1, min2, min3)

	const a int = 5
	const b int = 4

	const mult = a * b

	//const x int = 4
	//const pi float64 = 3.14
	//
	//const s = x * pi

	const x = 5 // untyped constant

	const i int = x     // engine does this i = int(x)
	const j float64 = x // engine does this j = float(x)
	const t byte = x

	fmt.Println(i, j, t) // prin
}
