package main

import (
	"fmt"
	"strings"
)

func main() {
	var employees map[string]string

	fmt.Printf("%#v\n", employees)
	fmt.Printf("Entries count %d\n", len(employees))

	// prints out an empty string for a non-existing key (zero-value)
	fmt.Printf("Value for key Dan is %s", employees["Dan"])

	// prints out 0.0000 for a non-existing key (zero-value)
	//var balances map[string]float64
	//fmt.Printf("Value for key Dan is %f\n", balances["Dan"])

	// can we use an array as a map key?
	// yes we can
	//var m1 map[[3]int]string
	//var employes map[string]string

	//employes["Dan"] = "Developer"

	// prints out
	// panic: assignment to entry in nil map
	// we cannot write to non-initialized map
	//fmt.Println(employes)

	people := map[string]string{}

	people["Kirill"] = "me"
	people["Anna"] = "she"

	fmt.Printf("%v\n", people)

	fmt.Println("\n" + strings.Repeat("@", 20))
	// declaring non-nil maps
	// where you can write
	// #1
	m1 := make(map[int]int)
	m1[4] = 5

	// #2
	m2 := map[int]int{2: 4, 5: 7}
	m2[2] = 9

	balances := map[string]float64{
		"RUB": 3.,
		"BTC": 0.001,
	}

	// writes non-existent key
	balances["USDT"] = 15.
	// updates existent key
	balances["RUB"] = 100.

	// returns zero-value for non-existing key
	solanaBalance := balances["SOL"]

	// prints map[BTC:0.001 RUB:100 USDT:15]
	fmt.Println(balances)
	// prints Sol balance is 0.000000, type float64
	fmt.Printf("Sol balance is %f, type %T", solanaBalance, solanaBalance)

	fmt.Println("\n" + strings.Repeat("@", 20))

	rates := map[string]float64{
		"RUB": 1.4,
		"BTC": 2.3,
		"TRK": 3.2,
	}

	rub, getR := rates["RUB"]
	usd, getUsd := rates["USD"]

	// prints out Value 1.4, is in map true
	fmt.Printf("Value %v, is in map %v\n", rub, getR)
	// prints out Value 0, is in map false
	fmt.Printf("Value %v, is in map %v\n", usd, getUsd)

	for k, v := range rates {
		fmt.Printf("Key: %s Val: %f\n", k, v)
	}

	// delete a key
	delete(rates, "RUB")
	fmt.Println(rates)

	fmt.Println("\n" + strings.Repeat("@", 20))
	compareMaps()

	fmt.Println("\n" + strings.Repeat("@", 20))
	cloningMap()
}

func compareMaps() {

	// maps comparison
	m1 := map[string]int{"A": 5}
	m2 := map[string]int{"A": 5}

	s1 := fmt.Sprintf("%s", m1)
	s2 := fmt.Sprintf("%s", m2)

	// prints out Maps are equal: true
	fmt.Printf("Maps are equal: %v\n", s1 == s2)
}

func cloningMap() {
	friends := map[string]int{"Dan": 40, "Kirill": 50}

	neighbours := friends

	neighbours["Dan"] = 999

	// prints out map[Dan:999 Kirill:50]
	fmt.Println(friends)

	fmt.Println("\n" + strings.Repeat("@", 20))

	// cloning a map
	people := map[string]int{"Dan": 40, "Kirill": 30}
	mates := make(map[string]int)

	for k, v := range people {
		mates[k] = v
	}

	mates["Dan"] = 0

	// prints out map[Dan:40 Kirill:30]
	fmt.Println(people)
	// prints out map[Dan:0 Kirill:30]
	fmt.Println(mates)

	employees := map[string]int{}
	// prints out 0
	fmt.Println(employees["Andrei"])

}
