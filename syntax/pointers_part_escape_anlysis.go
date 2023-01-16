package syntax

type user struct {
	name  string
	email string
}

func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@email.com",
	}
	println("V1", &u)
	return u
}

// u is created at heap amazingly
func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@email.com",
	}
	println("V2", &u)
	return &u
}

func PointersEscapAnalysis() {

	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", &u2)
}
