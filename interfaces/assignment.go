package main

import (
	"fmt"
	"strings"
)

type money float64

type book struct {
	price float64
	title string
}

func (m money) print() {
	fmt.Println(m.printStr())
}

func (m money) printStr() string {
	return fmt.Sprintf("%.2f", m)
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() {
	b.price = b.price * 0.9
}

func main() {
	var m = money(3.43334)
	m.print()

	fmt.Println("pr str result", m.printStr())

	fmt.Println("\n" + strings.Repeat("@", 20))

	b := book{price: 100, title: "Everyone has a story"}

	fmt.Println(b.vat())

	b.discount()
	fmt.Println(b)

	fmt.Println("\n" + strings.Repeat("@", 20))

	var v vehicle

	v = car{brand: "BMW", licenceNo: "33AAA99"}

	fmt.Println(v.License(), v.Name())

	var empty interface{}

	empty = 3
	empty = "sldkjf"
	empty = float64(4.4)
	empty = []int{2, 3, 4, 45}

	empty = append(empty.([]int), 2, 3)

	fmt.Printf("%v, %T\n", empty, empty)

}

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}
