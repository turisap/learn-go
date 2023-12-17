package main

import (
	"fmt"
	"time"
)

func main() {

	ch2 := make(chan int, 2) // BUFFERED

	go func(c chan int) {

		for i := 0; i < 10; i++ {
			fmt.Println("Before pushing to the channel", i)
			c <- i
			fmt.Println("After pushing to the channel", i)
		}

		close(c)
	}(ch2)

	fmt.Println("Sleep for 2 seconds")
	time.Sleep(2 * time.Second)

	for v := range ch2 {
		fmt.Printf("Received a message from the channel %d\n", v)
	}

	/* will print

	Before pushing to the channel 0
	After pushing to the channel 0
	Before pushing to the channel 1
	After pushing to the channel 1
	Before pushing to the channel 2
	After pushing to the channel 2
	Received a message from the channel 0
	Received a message from the channel 1
	Received a message from the channel 2
	*/

	time.Sleep(1 * time.Second)
}
