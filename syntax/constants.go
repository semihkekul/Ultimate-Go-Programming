package syntax

func Constants() {
	const ui = 12345
	const uf = 3.141592

	const ti int = 12345
	const tf float64 = 3.141592

	const myUint8 uint8 = 100

	// var a int = 3
	// var b float64 = 2.3
	// var c = a * b  // invalid operation: a * b (mismatched types int and float64)
	var c = ui * uf // works
	println(c)

	const (
		A = iota + 7
		B
		C
	)

	println(A, B, C)
}
