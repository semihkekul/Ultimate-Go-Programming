package syntax

import "fmt"

type example struct {
	flag    bool
	counter int16
	pi      float32
}

type bill struct {
	flag    bool
	counter int16
	pi      float32
}

type alice struct {
	flag    bool
	counter int16
	pi      float32
}

func StructTypes() {
	var e1 example
	fmt.Printf("%+v\n", e1)

	e2 := example{
		flag:    true,
		counter: 20,
		pi:      3.141592,
	}

	fmt.Printf("%+v\n", e2)

	// struct literal; anonymous type

	ls := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 40,
		pi:      6.02,
	}

	fmt.Printf("%+v\n", ls)

	var b bill
	var a alice
	var c alice

	b = bill(a)

	c = ls

	fmt.Println(a, b, c)
}
