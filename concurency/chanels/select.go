package main

import (
	"fmt"
	"time"
)

func r2(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello"
}

func r1(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "Salut"
}

func main() {
	fmt.Println("Start exec")
	const mil = 100000
	start := time.Now().UnixNano() / mil
	ch1 := make(chan string)
	ch2 := make(chan string)

	go r1(ch1)
	go r2(ch2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Message from ch1", msg1)

		case msg2 := <-ch2:
			fmt.Println("Message from the ch2", msg2)

		default:
			fmt.Println("no activity")
		}
	}

	fmt.Println("Finish Exec")
	end := time.Now().UnixNano() / mil
	fmt.Println(end - start)
}
