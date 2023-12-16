package main

import "fmt"

func main() {
	// chanel declaration (nil value)
	var ch chan int
	ch = make(chan int)

	// prints out 0x14000120000 it is like
	// a pointer
	fmt.Println(ch)

	// SEND operation
	ch <- 10

	// RECEIVE
	num := <-ch

	// CLOSE
	close(ch)

	_ = num

}
