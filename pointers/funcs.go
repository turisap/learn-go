package main

import (
	"fmt"
	"strings"
)

func change(a *int) {
	*a = 100
}

func changeVar(a int) {
	a = 66
}

func main() {
	x := 5
	y := 10
	p := &x

	// X before calling change 5
	fmt.Println("X before calling change", x)
	change(p)
	// X after calling change 100
	fmt.Println("X after calling change", x)

	// Y before calling change 10
	fmt.Println("Y before calling changeVar", y)
	changeVar(y)
	// Y before calling change 10
	fmt.Println("Y after calling changeVar", y)

	fmt.Println("\n" + strings.Repeat("@", 20))

}
