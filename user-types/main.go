package main

func main() {
	type speed uint

	var x uint = 5

	// custom type conversion
	var s3 = speed(x)

	_ = s3
}
