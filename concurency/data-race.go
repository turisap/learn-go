package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	t1 := time.Now()

	const gr = 100
	var wg sync.WaitGroup

	wg.Add(gr * 2)

	var n int = 0
	var m sync.Mutex

	for i := 0; i < 100; i++ {
		// Mutexes
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			// another option
			m.Lock()
			defer m.Unlock()
			n--
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("The final n value is: ", n)
	fmt.Println("Time finished", time.Now().Sub(t1).Milliseconds(), "ms")
}
