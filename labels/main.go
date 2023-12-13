package main

import (
	"fmt"
)

func main() {
	people := [5]string{"mark", "brand", "kirill", "ann", "lol"}
	freinds := [2]string{"mark", "ann"}
	outer := 21

	// @COOL
outer:
	for _, name := range people {
		for idx, fr := range freinds {
			if name == fr {
				fmt.Printf("Found a friend %s at index %d\n", fr, idx)
				break outer
			}
		}
	}

	fmt.Printf("Next instruction")
	fmt.Println(outer)
}
