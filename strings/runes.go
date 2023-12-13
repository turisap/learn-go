package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	a, b := 'A', 'B'

	fmt.Printf("Type of a and b is %T %T, values are %d, %d \n", a, b, a, b)

	// this is one Thai char
	// prints out 3
	// len return the number
	// of bytes, not runes
	str := "แ"
	fmt.Println(len(str))

	thaiCat := "แมว"

	// prints out à¹à¸¡à¸§
	// because it goes byte by byte
	// but one rune is more than one byte
	fmt.Println("\n" + strings.Repeat("@", 20))
	for i := 0; i < len(thaiCat); i++ {
		fmt.Printf("%c", thaiCat[i])
	}

	fmt.Println("\n" + strings.Repeat("@", 20))

	thaiCar := "รถ"
	// iterating char by char
	for i := 0; i < len(thaiCar); {
		rune, size := utf8.DecodeRuneInString(thaiCar[i:])
		i += size

		fmt.Printf("%c", rune)
	}

	fmt.Println("\n" + strings.Repeat("@", 20))

	thaiHouse := "บ้าน"
	// iterating char by char #2
	for _, rune := range thaiHouse {
		fmt.Printf("%c", rune)
	}

	fmt.Println("\n" + strings.Repeat("@", 20))

	// runes count in a string
	// prints Chars in string 11, bytes in string 33
	thaiPhuket := "เมืองภูเก็ต"
	c := utf8.RuneCountInString(thaiPhuket) // @COOL
	fmt.Printf("Chars in string %d, bytes in string %d", c, len(thaiPhuket))
}
