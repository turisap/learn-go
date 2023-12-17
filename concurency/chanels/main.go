package main

import "fmt"

func f1(n int, ch chan int) {
	ch <- n + 1
}

func main() {
	// two ways
	ch1 := make(chan int)

	// only to receive
	//ch2 := make(<-chan string)

	// only to send
	//ch3 := make(chan<- string)

	go f1(10, ch1)

	n := <-ch1

	fmt.Println("Recieved value: ", n)
	fmt.Println("Exiting main")
}
