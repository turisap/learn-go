package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var r = 'ă'
	var s1 = "ma"
	var s2 = "m"

	fmt.Printf("%T\n", r)

	s2 = s2 + string(r)
	res := strings.Join([]string{s1, s2}, "")

	fmt.Printf(res)

	fmt.Println("\n" + strings.Repeat("@", 20))

	s3 := "țară means country in Romanian"

	for i := 0; i < len(s3); i++ {
		fmt.Println(string(s3[i]))
	}

	fmt.Println("\n" + strings.Repeat("@", 20))

	for _, r = range s3 {
		fmt.Println(string(r))
	}

	fmt.Println("\n" + strings.Repeat("@", 20))

	for i := 0; i < len(s3); {
		r, size := utf8.DecodeRuneInString(s3[i:])
		i += size
		fmt.Println(string(r))
	}

	fmt.Println("\n" + strings.Repeat("3", 20))

	s4 := "Go is cool!"
	x := s1[0]
	fmt.Println(x)

	s4 = strings.Replace(s4, "G", "x", 1)
	fmt.Println(s4)

	// printing the number of runes in the string
	fmt.Println(utf8.RuneCountInString(s4))

	fmt.Println("\n" + strings.Repeat("4", 20))

	// convert a string into a byteSlice
	s5 := "你好 Go!"
	byteSlice := []byte(s5)

	fmt.Println(byteSlice)

	for idx, v := range byteSlice {
		fmt.Printf("Byte %d as index %d\n", v, idx)
	}

	fmt.Println("\n" + strings.Repeat("5", 20))

	s6 := "你好 Go!"
	runeSlice := []rune(s6)

	for idx, v := range runeSlice {
		fmt.Printf("Rune %c at index %d\n", v, idx)
	}
}
