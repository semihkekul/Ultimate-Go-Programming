package syntax

func PointersPassByValue() {
	count := 10

	println("count:\t Value of [", count, "]\tAddr of[", &count, "]")

	increment(count)

	println("count:\t Value of [", count, "]\tAddr of[", &count, "]")

}

func increment(inc int) {
	inc++
	println("inc:\t Value of [", inc, "]\tAddr of[", &inc, "]")
}
