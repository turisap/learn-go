package main

import "math"

func main() {
	var x uint8 = 255

	x++
	x++

	// prints out 1 (overflow starts back from 0)
	println(x)

	//var b = uint8(255 + 1)

	var b int8 = -128
	b--

	// prints out +127
	println(b)

	var f float32 = float32(math.MaxFloat32)
	println(f)

	f *= 1.2

	// prints +Inf. Floats overflow to Infinity
	println(f)
}
