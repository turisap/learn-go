package main

func main() {

	// DEFINED type
	type second uint8

	// ALIAS type
	type tree = string

	var a uint8 = 8
	var b byte = 1

	a = b // no type mismath error

	_ = a
}
