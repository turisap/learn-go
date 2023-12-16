package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("F1 execution started")
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("F1 execution ended")
	wg.Done()
}

func f2() {
	fmt.Println("F2 execution started")
	for i := 10; i <= 13; i++ {
		fmt.Println(i)
	}
	fmt.Println("F2 execution ended")
}

func systemInfo() {
	fmt.Println("CPUs number", runtime.NumCPU())
	fmt.Println("Goroutines number", runtime.NumGoroutine())
	fmt.Println("ENV", runtime.GOARCH, runtime.GOOS)

	// used to control how many CPU cores
	// go can use
}
func main() {
	fmt.Println("main execution started")

	var wg sync.WaitGroup
	wg.Add(1)

	go f1(&wg)
	fmt.Println("Goroutines number", runtime.NumGoroutine())

	f2()

	wg.Wait()
	fmt.Println("main execution ended")
}
