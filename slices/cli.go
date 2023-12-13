package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var numbers = os.Args[1:]

	if len(numbers) < 2 || len(numbers) > 10 {
		fmt.Println("The program accepts between 2 and 10 arguments")

		return
	}

	var s int64 = 0
	var p int64 = 1

parser:
	for _, v := range numbers {
		val, err := strconv.ParseInt(v, 10, 32)

		if err != nil {
			fmt.Printf("Cannot parse input %s", v)
			break parser
		}

		s += val
		p *= val
	}

	fmt.Printf("Sum %d; Product: %d\n", s, p)
}
