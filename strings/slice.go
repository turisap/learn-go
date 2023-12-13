package main

import "fmt"

func main() {
	// slicing ASCI string
	s1 := "I love Golang"

	// prints out "lov"
	fmt.Println(s1[2:5])

	// prints out �
	// which is Unicode representation
	// of the first byte
	thaiGirl := "สาว"
	fmt.Println(thaiGirl[0:1])

	// getting a slice of runes out of string
	// prints out "中文"
	s2 := "中文维基是世界上"
	slice := []rune(s2)
	fmt.Println(string(slice[0:2]))

	// prints out 72 which is
	// a Unicode code point
	k := "Go is cool!"
	fmt.Println(k[0])
}
