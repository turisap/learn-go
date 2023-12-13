package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "I love programming"

	// true
	fmt.Println(strings.Contains(s1, "love"))

	// true
	fmt.Println(strings.ContainsAny(s1, "xyl"))

	// false
	fmt.Println(strings.ContainsAny(s1, "xy"))

	fmt.Println(strings.Count("cheese", "g"))

	// 5
	fmt.Println(strings.Count("Five", ""))

	// 2 @COOL
	fmt.Println(strings.Count("Fivev", "v"))

	// java, python
	fmt.Println(strings.ToLower("JaVa, PYThoN"))

	// JAVA, PYTHON
	fmt.Println(strings.ToUpper("JaVa, PYThoN"))

	// true
	fmt.Println("go" == "go")

	// false
	fmt.Println("GO" == "go")

	// compares two strings case-insensitive
	// prints out true @COOL
	fmt.Println(strings.EqualFold("Go", "gO"))

	// repeat string
	fmt.Println("\n" + strings.Repeat("@", 20))

	// replace N occurrences (-1 is the same as replace All)
	s2 := strings.Replace("123.12.56.0", ".", ":", -1)
	fmt.Println(s2)

	// LULULIPUP
	s3 := strings.ReplaceAll("LOLOLIPOP", "O", "U")
	fmt.Println(s3)

	fmt.Println("\n" + strings.Repeat("@", 20))
	s4 := strings.Split("a,b,c,d", ",")
	fmt.Println(s4)

	s5 := strings.Split("Go forest go", "")
	fmt.Println(s5)

	s6 := strings.Split("ฉันกำ", "")
	// prints out [ฉ ั น ก ำ]
	fmt.Println(s6)

	sl := []string{"I", "love", "dating"}
	fmt.Printf(strings.Join(sl, "") + "\n")
	fmt.Printf(strings.Join(sl, " ") + "\n")

	fmt.Println("\n" + strings.Repeat("@", 20))

	fmt.Println(strings.Trim("    Love dkfj      df", " "))
	fmt.Println(strings.Trim("??KIRILL YOU??", "?"))
	fmt.Println(strings.Trim("////Love dkfj      df///", "/"))
	fmt.Printf(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))

	fmt.Println("\n" + strings.Repeat("@", 20))

	// prints out []string not []rune
	fmt.Printf("%T", strings.Split("KITILL", ""))
}
