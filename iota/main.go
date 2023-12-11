package main

func main() {
	const (
		a = iota
		b = iota
		c = iota
	)

	println(a, b, c)

	const (
		j = -(iota + 2)
		i
		l
	)

	println(j, i, l)
}
