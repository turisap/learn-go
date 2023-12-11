package main

import (
	"fmt"
)

func main() {
	// THIS WAY
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	//// OR
	//for j := 100; j > 80; {
	//	fmt.Println(j)
	//	j--
	//}
	//
	//// THIS IS WHILE LOOP analog
	//k := 10
	//for k < 20 {
	//	fmt.Printf("K is %d\n", k)
	//	k++
	//}

	//for s := 10; s < 20; s++ {
	//	if s%2 != 0 {
	//		continue
	//	}
	//
	//	fmt.Printf("Value is %v\n", s)
	//}

	//for s := 10; s < 20; s++ {
	//	if s == 17 {
	//		break
	//	}
	//
	//	fmt.Printf("Value is %v\n", s)
	//}

	// WHILE TRUE ANALOG
	count := 0
	for i := 0; true; i++ {

		if i%3 == 0 {
			count++
		}

		if count > 10 {
			fmt.Printf("i is %d and count is %d \n", i, count)
			break
		}
	}
}
