package main

import (
	"fmt"
	"strings"
)

type names []string

func (n names) print() {
	for i, name := range n {
		fmt.Printf("%s at %d\n", strings.ToUpper(name), i)
	}
}

/* pointer receiver */
type car struct {
	brand string
	price int
}

func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c *car) changeCarP(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
}

func main() {
	//n := names{"kirill", "anna"}

	//n.print()

	c := car{brand: "BMW", price: 344}

	c.changeCar1("KAMAZ", 0)

	fmt.Println(c)

	c.changeCarP("KAMAZ", 0)

	fmt.Println(c)

	fmt.Println("\n" + strings.Repeat("@", 20))
	var yourCar *car
	yourCar = &c

	fmt.Println(*yourCar)

	// call methods on pointers
	yourCar.changeCarP("LOL", 33)

	fmt.Println(yourCar)

	(*yourCar).changeCarP("KIS", 66)

	fmt.Println(*yourCar)
	fmt.Println(c)

}
