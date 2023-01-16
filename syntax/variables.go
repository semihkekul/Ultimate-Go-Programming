package syntax

import "fmt"

func Variables() {
	// zero value init !!!
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int\t %T [%v]\n", a, a)
	fmt.Printf("var b string\t %T [%v]\n", b, b)
	fmt.Printf("var c float64\t %T [%v]\n", c, c)
	fmt.Printf("var b bool\t %T [%v]\n", d, d)

	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10\t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\"\t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159\t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true\t %T [%v]\n", dd, dd)

	// no casting -> conversion
	aaa := int32(10)

	fmt.Printf("aaa := int32(10)\t %T [%v]\n", aaa, aaa)
}
