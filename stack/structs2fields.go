package stack

// S2F is a 2 field struct.
type S2F struct {
	A string `json:"a"`
	B string `json:"b"`
}

// S2fByValue is method to test passing structs
func S2fByValue() {
	x := S2F{A: "naruto", B: "rocks"}
	s2fByValue2(x)
}

func s2fByValue2(x S2F) S2F {
	x.A = "not naruto"
	return s2fByValue3(x)
}

func s2fByValue3(x S2F) S2F {
	x.B = "not rocks"
	return x
}

// S2fByPointer is method to test passing pointers
func S2fByPointer() {
	x := S2F{A: "naruto", B: "rocks"}
	s2fByPointer2(&x)
}

func s2fByPointer2(x *S2F) *S2F {
	x.A = "not naruto"
	return s2fByPointer3(x)
}

func s2fByPointer3(x *S2F) *S2F {
	x.B = "not rocks"
	return x
}
