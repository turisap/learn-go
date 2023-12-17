package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func sum(f1 float64, f2 float64, wg *sync.WaitGroup) {

	s := fmt.Sprintf("%.2f", f1+f2)
	fmt.Println(s)

	defer wg.Done()
}

func ex1() {
	fmt.Println("Start exec")
	var wg sync.WaitGroup
	l := 10

	wg.Add(l)

	for i := 0; i < l; i++ {
		go sum(float64(i)+0.5, float64(i)+0.3, &wg)
	}

	wg.Wait()
	fmt.Println("Finish exec")
}

func ex2() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func(f float64, s *sync.WaitGroup) {
		defer s.Done()
		time.Sleep(3 * time.Second)
		fmt.Println(math.Sqrt(f))
	}(55.5, &wg)

	wg.Wait()
}

func ex3() {
	var wg sync.WaitGroup

	wg.Add(50)
	for i := 100; i < 150; i++ {
		go func(f float64, s *sync.WaitGroup) {
			defer s.Done()
			fmt.Println(math.Sqrt(f))
		}(float64(i), &wg)
	}
	wg.Wait()
}

func deposit(b *int, n int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()

	*b += n

	mu.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, mu *sync.Mutex) {

	mu.Lock()

	*b -= n

	mu.Unlock()
	wg.Done()
}

func ex4() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &mu)
		go withdraw(&balance, i, &wg, &mu)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}

func ex5() {
	var ch1 chan float64
	var ch2 = make(chan<- float64)
	var ch3 = make(<-chan float64)
	var ch4 = make(chan float64, 10)

	fmt.Printf("%T\n", ch1)
	fmt.Printf("%T\n", ch2)
	fmt.Printf("%T\n", ch3)
	fmt.Printf("%T\n", ch4)
}

func ex6() {
	fmt.Println("start execution")
	ch := make(chan string)

	go func(s string, c chan string) {
		fmt.Println("sending the message")
		time.Sleep(2 * time.Second)
		c <- s
		fmt.Println("sent the message")
	}("HELLO", ch)

	fmt.Println(<-ch)
	fmt.Println("Stop execution")
}

func ex7() {
	c := make(chan int)

	go func(n int) {
		c <- n
	}(100)

	fmt.Println(<-c)
}

func power(x int, ch chan float64) {
	res := math.Pow(float64(x), 2)

	fmt.Printf("sending %f to the result back\n", res)
	ch <- res
}

func main() {
	ex7()
}
