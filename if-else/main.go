package main

func main() {
	var price, inStock = 77, true

	if price > 80 {
		println("Too expensive")
	} else if price < 80 && inStock {
		println("Buy it")
	}
}
