package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := []int{2, 3, 4, 5}

	mightMutate(nums...)

	// prints out [999 3 4 5]
	// is data being mutated in mightMutate?
	fmt.Println(nums)

	s, p := SumAndProduct(1., 2., 3., 4.0)

	fmt.Printf("%f, %f\n", s, p)

	k := personInfo(34, "kirill", "shakirov", "turisap")
	fmt.Println(k)
}

func mightMutate(args ...int) {
	args[0] = 999
}

// variadic functions
func vFn(args ...int) {
	fmt.Printf("%#v\n", args)
}

func SumAndProduct(args ...float64) (float64, float64) {
	s := 0.
	p := 1.

	for _, v := range args {
		s += v
		p *= v
	}

	return s, p
}

func personInfo(age int, names ...string) string {
	fullName := strings.Join(names, " ")

	return fmt.Sprintf("%s of age %d", fullName, age)
}
