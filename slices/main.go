package main

import (
	"fmt"
)

func main() {
	var s []int

	// prints out  type is []int
	fmt.Printf("values:type is %T\n", s)
	// zero value  of non-initialized slice is nil
	// prints out true
	fmt.Println("s is equal to nil", s == nil)

	nums := []int{2, 3, 4, 5}

	fmt.Println(nums)

	// @COOL
	// init a slice with make
	var s1 = make([]int, 4)

	// prints out value is [0,0,0,0]
	fmt.Printf("Value is %v\n", s1)

	s2 := []int{1, 2, 3, 4, 5}
	for idx, v := range s2 {
		fmt.Println(idx, v)
	}

	sliceCompare()
	sliceAppend()
	sliceCopy()
	sliceExpressions()
}

func sliceCompare() {
	var n []int

	// print true
	fmt.Println("N is nil", n == nil)

	// print false
	var m = []int{}
	fmt.Println("M in nil", m == nil)

	a, b := []int{1, 2, 5}, []int{1, 2, 3}
	var eq bool = true

mainFn:
	for idx, v := range a {
		if v != b[idx] {
			eq = false
			break mainFn
		}
	}

	if eq {
		fmt.Println("Slices are equal")
	} else {
		fmt.Println("Slices are not equal")
	}
}

func sliceAppend() {
	// appending to a slice
	var s1 = []int{1, 2, 3}

	s1 = append(s1, 10, 111, 12)

	fmt.Println(s1)

	var s2 = []int{99}

	// appending using ellipsis operator
	var s3 = append(s2, s1...)

	fmt.Println(s3)
}

func sliceCopy() {
	// copy with copy
	src := []int{1, 2, 3}
	dist := make([]int, len(src))
	nn := copy(dist, src)

	fmt.Printf("Copy %d elements from src to dist. Resulting slice is %v\n", nn, dist)
}

func sliceExpressions() {
	// array
	a := [5]int{1, 2, 3, 4, 5}

	s := a[1:3]
	// prints out Sliced array is [2 3], type is []int (slice)
	fmt.Printf("Sliced array is %v, type is %T\n", s, s)

	b := [6]int{1, 2, 3, 4, 5, 6}
	// @COOL
	s1 := b[2:]
	// prints out [3 4 5 6]
	fmt.Println(s1)

	// @COOL
	s2 := b[:2]
	fmt.Println(s2)

	// @COOL
	s3 := b[:]
	fmt.Println(s3)

	// overwrite the last element
	k := [3]int{1, 2, 3}
	s4 := append(k[:len(k)-1], 100)
	fmt.Println(s4)
}
