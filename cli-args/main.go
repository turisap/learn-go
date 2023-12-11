package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Os ARG 1", os.Args[0])
	fmt.Println("Os ARG 2", os.Args[1])
	fmt.Println("Os ARG length", len(os.Args))

	var amount, err = strconv.ParseFloat(os.Args[1], 64)

	fmt.Printf("Type of os.Arg[1] is %T\n", os.Args[1])
	fmt.Printf("Type of is %T and value is %f\n", amount, amount)

	_ = err
}
