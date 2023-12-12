package main

import "fmt"

func main() {
	n := 5

	switch true {
	case n%2 == 0:
		fmt.Printf("Even n is %d\n", n)
	case n%2 != 0:
		fmt.Printf("Odd n is %d\n", n)
	}
}
