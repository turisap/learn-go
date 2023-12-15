package main

import (
	"fmt"
	"strings"
)

func main() {
	var x float64 = 4.5
	p1 := &x
	pp1 := &p1

	**pp1++

	// prints 5.5
	fmt.Println(**pp1)

	fmt.Println(x)

	fmt.Printf("p1 val %p, p1 address %p\n", p1, &p1)
	fmt.Printf("pp1 val %p, p2 address %p\n", pp1, &pp1)

	fmt.Println("\n" + strings.Repeat("@", 20))

	var p3 *int
	var p4 *int
	var p5 *int

	// (*int)(nil)
	fmt.Printf("%#v\n", p3)
	// prints true because
	// both are nil
	fmt.Println(p3 == p4)

	fmt.Println("\n" + strings.Repeat("@", 20))

	y := 4
	z := 4

	p3 = &y
	p4 = &z

	p5 = &y

	// false (points to different vars)
	fmt.Println(p3 == p4)
	// true (points to same        vars)
	fmt.Println(p3 == p5)

}
