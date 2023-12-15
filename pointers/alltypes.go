package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool, nums [3]int) {
	quantity = 3
	price = 500.5
	name = "Mobile Phone"
	sold = false
	nums[0] = 999999
}

func changeValueByPointer(quantity *int, price *float64, name *string, sold *bool) {

	*quantity = 3
	*price = 500.5
	*name = "Mobile Phone"
	*sold = false
}

type Product struct {
	price float64
	name  string
}

func productDiscount(p Product) {
	p.price = 0.5
	p.name = "OLOLO"
}

// p is a struct
func productDiscountPointer(p *Product) {
	(*p).price = 0.5 // this
	p.name = "OLOLO" // and this are the same
}

func changeSlice(s []int) {
	for i, _ := range s {
		s[i] *= 3
	}
}

func changeMap(m map[string]int) {
	m["a"] = 20
	m["x"] = 40
	m["b"] = 60
}

func main() {
	quantity, price, name, sold := 5, 300.2, "Laptop", true
	nums := [3]int{1, 2, 3}

	//BEFORE calling changeValues(): 5 300.2 Laptop true [1,2,3]
	fmt.Println("BEFORE calling changeValues():", quantity, price, name, sold, nums)

	changeValues(quantity, price, name, sold, nums)
	//AFTER calling changeValues(): 5 300.2 Laptop true [1,2,3]
	fmt.Println("AFTER calling changeValues():", quantity, price, name, sold, nums)

	//BEFORE calling changeValuesByPointer(): 5 300.2 Laptop true
	fmt.Println("BEFORE calling changeValues():", quantity, price, name, sold)

	changeValueByPointer(&quantity, &price, &name, &sold)
	//AFTER calling changeValuesByPointer():  3 500.5 Mobile Phone false
	fmt.Println("AFTER calling changeValues():", quantity, price, name, sold)

	/* STRUCTS */

	gift := Product{price: 100.0, name: "bicycle"}

	fmt.Printf("The gift before the change %v\n", gift)

	productDiscount(gift)

	fmt.Printf("The gift after the change %v\n", gift)

	// The gift before the change {100 bicycle}
	fmt.Printf("The gift before the change %v\n", gift)

	productDiscountPointer(&gift)

	// The gift after the change {0.5 OLOLO}
	fmt.Printf("The gift after the change %v\n", gift)

	s := []int{1, 2, 3}

	// Slice before change calling changeSlice() [1 2 3]
	fmt.Printf("Slice before change calling changeSlice() %v\n", s)

	changeSlice(s)

	// Slice after change calling changeSlice() [3 6 9]
	fmt.Printf("Slice after change calling changeSlice() %v\n", s)

	m := map[string]int{
		"K": 2,
		"r": 3,
	}

	fmt.Printf("Map before change, %v\n", m)

	changeMap(m)

	fmt.Printf("Map after change, %v\n", m)
}
