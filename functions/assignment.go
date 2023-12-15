package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(cube(3))
	fmt.Println(f1(5))
	fmt.Println(f2("5"))
	fmt.Println(f3(1, 2, 3, 4, 5))
	fmt.Println(searchItem([]string{"lion", "tiger", "bear"}, "tiger"))
	fmt.Println(searchItem([]string{"lion", "tiger", "bear"}, "panda"))
	fmt.Println(searchItemI([]string{"lion", "tiger", "bear"}, "panda"))
	fmt.Println(searchItemI([]string{"lion", "tiger", "bear"}, "TiGer"))

	customPrint := print

	customPrint(3)
	customPrint(5)
}

func cube(n float64) float64 {
	return math.Pow(n, 3.0)
}

func f1(n uint) (int, int) {
	var factorial = 1
	var i = 1
	var sum = 0

	for i = 1; i <= int(n); i++ {
		factorial = factorial * i
		sum += i
	}

	return factorial, sum
}

func f2(n string) int {
	n1, _ := strconv.Atoi(n)

	n2, _ := strconv.Atoi(n + n)

	n3, _ := strconv.Atoi(n + n + n)

	return n1 + n2 + n3
}

func f3(args ...int) (sum int) {

	for _, v := range args {
		sum += v
	}

	return
}

func searchItem(sl []string, s string) bool {

	for _, v := range sl {
		if v == s {
			return true
		}
	}

	return false
}
func searchItemI(sl []string, s string) bool {

	for _, v := range sl {
		if strings.EqualFold(v, s) {
			return true
		}
	}

	return false
}

func print(n int) {
	fmt.Println(n)
}
