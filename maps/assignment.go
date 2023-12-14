package main

import "fmt"

func main() {
	e1()
	e2()
}

func e1() {
	// panic: assignment to entry in nil map
	//var m1 map[int]bool
	//m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	s2 := fmt.Sprintf("%s", m2)
	s3 := fmt.Sprintf("%s", m3)

	fmt.Println(s2 == s3)
}

func e2() {
	m := map[int]bool{1: true, 2: false, 3: false}

	delete(m, 1)

	for k, v := range m {
		fmt.Printf("Key: %d, Value: %v\n", k, v)
	}
}
