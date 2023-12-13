package main

import "fmt"

func main() {
	var arr [4]int
	// zero-initialized to {0,0,0,0}
	fmt.Printf("values: %v\n", arr)
	//fmt.Printf("values and type: %#v\n", arr)

	var a1 = [4]int{1, 2, 3, 4}
	fmt.Printf("values: %#v\n", a1)

	var a2 = [3]string{"Kokia", "ti", "mar"}
	fmt.Printf("values: %#v\n", a2)

	var a3 = [4]int{1, 2}
	fmt.Printf("values: %#v\n", a3)

	// partlial values
	var a4 = [4]int{1, 2}
	fmt.Printf("values: %#v\n", a4)

	// allows put a dynamic number of elements
	// upon initialization
	var a5 = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("values: %#v\n", a5)
	fmt.Printf("length: %d\n", len(a5))

	fmt.Println("=====================")

	var aN = [4]int{1, 2, 3, 4}
	fmt.Printf("Value is %d\n", aN[0])

	aN[0] = 999
	fmt.Printf("Value is %d\n", aN[0])

	// two ways to iterate over an array
	for i, v := range aN {
		fmt.Printf("Value is %d idx: %d\n", v, i)
	}

	for i := 0; i < len(aN); i++ {
		fmt.Printf("Value is %d idx: %d\n", aN[i], i)
	}

	balancies := [2][3]int{
		{5, 6, 7},
		[3]int{3, 4, 5},
	}

	fmt.Println(balancies)

	m := [3]int{1, 2, 3}
	// does a deep copy (new array)
	n := m

	// prints true
	fmt.Printf("M is equal to N: %v\n", m == n)

	m[0] = -1
	// prints false
	fmt.Printf("M is equal to N: %v\n", m == n)

	// indexed arrays

	grades := [3]int{
		0: 1,
		2: 5,
		1: 4,
	}

	fmt.Println(grades)

	bal := [3]int{2: 50}

	// prints [0 0 50]
	fmt.Println(bal)

	teachers := [...]string{5: "dan"}
	// prints [     dan] 6 -> 5 empty strings and "dan"
	fmt.Println(teachers, len(teachers))

	cities := [...]string{
		5: "London",
		// unkeyd element gets its index
		// from the last keyd one
		// prints out  [ Moscow    London Kemerovo]
		"Kemerovo",
		1: "Moscow",
	}

	fmt.Println(cities)

	checked := [...]bool{5: true, 6: true}

	// prints out
	// [false false false false false true true]
	fmt.Println(checked)

	nums1 := [...]int{10, 20, 30}
	nums2 := nums1
	nums2[0] = 1

	// prints [10 20 30]
	fmt.Println(nums1)

	ch()
}

func ch() {
	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10

	myArray[0] = float64(a)

	myArray[2] = 10.10

	fmt.Println(myArray)

}
