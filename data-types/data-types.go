package main

import "fmt"

func main() {
	var num int8 = 123

	fmt.Printf("%T\n", num)

	var i2 uint16 = 65530

	fmt.Printf("%T\n", i2)

	var f1, f2, f3 = 0.5, 1.4, 0.3

	fmt.Printf("%T, %T, %T\n", f1, f2, f3)

	fmt.Println("============")

	var char rune = 'f'
	//prints 102 (ASCI code for 'f')
	fmt.Printf("%d\n", char)

	var logged bool = true

	fmt.Printf("%T\n", logged)

	var s string = "hello go"
	fmt.Printf("%T\n", s)
	fmt.Printf("%s\n", s)

	fmt.Println("================= COMPOSITE ============")
}
