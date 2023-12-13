package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5}

	s2, s3 := s1[1:3], s1[2:4]

	s2[1] = 99999

	// will all of them contain 999?
	/*
		S1 [1 2 99999 4 5]
		S2 [2 99999]
		S3 [99999 4]
	*/
	fmt.Printf("S1 %v\n", s1)
	fmt.Printf("S2 %v\n", s2)
	fmt.Printf("S3 %v\n", s3)

	arr1 := [4]int{1, 2, 3, 4}

	s5, s6 := arr1[0:2], arr1[1:3]

	arr1[1] = 666

	// will all elements contain 666?
	/*
		[1 666]
		[666 3]
		[1 666 3 4]
	*/
	fmt.Println(s5)
	fmt.Println(s6)
	fmt.Println(arr1)

	differentSlices()
}

func differentSlices() {
	fmt.Println("===========")
	// two distinct slices
	// do not have one backing array
	cars := []string{"Nissan", "BWM", "KAMAZ"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	newCars[0] = "OLOLO"

	/*
		[Nissan BWM KAMAZ]
		[OLOLO BWM]
	*/
	fmt.Println(cars)
	fmt.Println(newCars)

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	newSlice := nums[0:2]

	// Slice length 2, slice capacity 7
	fmt.Printf("Slice length %d, slice capacity %d\n", len(newSlice), cap(newSlice))

	newSlice = nums[5:7]
	// Slice length 2, slice capacity 2 (because there are only 2el in the backing array after
	// idx of 5)
	fmt.Printf("Slice length %d, slice capacity %d\n", len(newSlice), cap(newSlice))

	s1 := []int{1, 2, 3, 4}
	s2 := s1[0:1]

	// prints addresses of
	// the slice headers
	fmt.Printf("%p\n", &s1)
	fmt.Printf("%p\n", &s2)

	arr := [5]int{1, 2, 3, 4, 5}
	sliceArr := []int{1, 2, 3, 4, 5}

	// which is of bigger size?
	// the array or the slice
	/*
		arr size is 40
		sliceArr size is 24
	*/
	fmt.Printf("arr size is %d\n", unsafe.Sizeof(arr))
	fmt.Printf("sliceArr size is %d\n", unsafe.Sizeof(sliceArr))
}
