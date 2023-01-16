package syntax

func PointersSharingData() {
	count := 33

	println("count:\t Value of [", count, "]\tAddr of[", &count, "]")

	increment2(&count)

	println("count:\t Value of [", count, "]\tAddr of[", &count, "]")
}

func increment2(inc *int) {
	*inc++
	println("*inc:\t Value of [", *inc, "]\tAddr of[", &inc, "]\tValue points to [", inc, "]")
}
