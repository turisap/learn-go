package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int) // unbuffered

	go func(c chan int) {
		fmt.Println("Before pushing to the channel") // 2
		// routine execution block here until someone
		// gets data from the chanel
		c <- 10
		fmt.Println("After pushing to the channel") // 5
	}(ch1)

	fmt.Println("Sleep for 2 seconds") // 1
	time.Sleep(2 * time.Second)

	fmt.Println("Main goroutines starts to receive the data") // 3
	d := <-ch1
	fmt.Println("Main goroutines received data D:", d) // 4
	time.Sleep(1 * time.Second)
}
