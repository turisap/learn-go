package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if i, err := strconv.Atoi("20"); err == nil {
		fmt.Println(i)
	} else {
		fmt.Println(err)
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("The program accepts only 1 argument")
	} else if i, err := strconv.Atoi(args[1]); err == nil {
		fmt.Printf("Type of i is %T and value is %d\n", i, i)
	} else {
		fmt.Printf("Cannot convert %s to integer\n", args[1])
	}
}
