package main

import "fmt"

func main() {
	age := 30
	var ololo = "kis"
	_ = "KII"

	println(age, ololo)

	s := "learn Go lang"
	fmt.Println(s)

	car, cost := "nissan", 4000

	fmt.Println(car, cost)

	car, year := "BMW", 2016

	fmt.Println(car, year)

	var opened = false

	opened, file := true, "kis"

	println(opened, file)

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 4, 5

	fmt.Println(i, j)

	i, j = j, i
	fmt.Println(i, j)
}
