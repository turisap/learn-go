package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	const gr = 100
	var wg sync.WaitGroup

	wg.Add(gr * 2)

	var n int = 1000

	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			n++
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("The final n value is: ", n)
}
