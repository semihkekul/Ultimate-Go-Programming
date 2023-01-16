package syntax

const size = 1024

func PointersStackGrowth() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}

func stackCopy(s *string, c int, a [size]int) {
	println(c, s, *s)
	c++
	if c == 10 {
		return
	}

	stackCopy(s, c, a)
}

// 0 0xc000143f50 HELLO
// 1 0xc000143f50 HELLO
// 2 0xc00015df50 HELLO stack is copied to a bigger memory
// 3 0xc00015df50 HELLO
// 4 0xc00015df50 HELLO
// 5 0xc00015df50 HELLO
// 6 0xc00019ff50 HELLO stack is copied to a bigger memory again
// 7 0xc00019ff50 HELLO
// 8 0xc00019ff50 HELLO
// 9 0xc00019ff50 HELLO
