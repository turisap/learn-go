package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := string(109)
	// prints out m
	// converts from ASCI code
	println(s)

	// converting float to a string
	var myStr = fmt.Sprintf("%f", 3.44)
	println(myStr)

	// converting integer to a string
	var myStrI = fmt.Sprintf("%d", 344)
	println(myStrI)

	// will print out ൴
	println(string(3444))

	println("==============")

	var str = "4.554234"
	// prints type of string
	fmt.Printf("Type of f is %T\n", str)

	var f, err = strconv.ParseFloat(str, 64)
	_ = err
	// prints type of float
	fmt.Printf("Type of f is %T\n", f)

	i, err := strconv.Atoi("50") // asci string to integer
	s1 := strconv.Itoa(20)       // integer to ASCI

	fmt.Printf("Type of  is %T\n value %d\n", i, i)
	fmt.Printf("Type of  is %T value %s\n", s1, s1)
	println(err)
}
