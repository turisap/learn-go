package main

import "fmt"

func main() {
	var nums = []int{}

	// Length: 0, Capacity: 0 -> backing array has length of 0,
	// which is fixed in go
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))

	// Length: 2, Capacity: 2
	// created a new backing array
	nums = append(nums, 1, 2)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))

	// Length: 3, Capacity: 4
	// crated another backing array of higher capacity
	nums = append(nums, 5)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))

	// Length: 5, Capacity: 8
	// crated another backing array of higher capacity
	nums = append(nums, 300, 500)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))

	fmt.Println("=================")
	letters := []string{"A", "B", "C", "D", "E", "F"}
	letters = append(letters[0:1], "X", "Y")

	// prints out [A X Y]
	fmt.Println(letters)
	// an error because there is no index of 4 in the slice
	//letters[4]

	// BUT the slice can see the whole backing array
	// prints out [D E F]
	fmt.Println(letters[3:6])

	n1 := []int{10, 20, 30, 40, 50}
	n2 := append(n1, 100)
	n2[0] = 66

	// prints out [10 20 30 40 50]
	// because there was a new backing array
	// created for n2
	fmt.Println(n1)

}
