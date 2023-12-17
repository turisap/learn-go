package main

import (
	"fmt"
	"strings"
)

func factorial(n int, c chan int) {
	f := 1

	for i := 2; i <= n; i++ {
		f *= i
	}

	c <- f
}

func main() {
	var ch = make(chan int)
	defer close(ch)

	// main blocks and put on sleep
	go factorial(5, ch)

	// when a message received, main unblocks
	f := <-ch

	fmt.Printf("Factorial of 5 is %d\n", f)

	//for i := 1; i < 30; i++ {
	//	go factorial(i, ch)
	//	f = <-ch
	//	fmt.Printf("Factorial of %d is %d\n", i, f)
	//}

	for i := 5; i < 15; i++ {

		go func(n int, channel chan int) {
			f := 1

			for i := 2; i <= n; i++ {
				f *= i
			}

			channel <- f
		}(i, ch)

		fmt.Printf("Factorial of %d is %d\n", i, <-ch)
		fmt.Println("\n" + strings.Repeat("@", 20))
	}
}
