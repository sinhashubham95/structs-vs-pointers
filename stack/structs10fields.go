package stack

// S10F is a 10 field struct.
type S10F struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
	E string `json:"e"`
	F string `json:"f"`
	G string `json:"g"`
	H string `json:"h"`
	I string `json:"i"`
	J string `json:"j"`
}

// S10fByValue is method to test passing structs
func S10fByValue() {
	x := S10F{A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto", F: "naruto", G: "rocks", H: "hinata", I: "rocks", J: "boruto"}
	s10fByValue2(x)
}

func s10fByValue2(x S10F) S10F {
	x.A = "not naruto"
	x.B = "not rocks"
	x.F = "not naruto"
	x.G = "not rocks"
	return s10fByValue3(x)
}

func s10fByValue3(x S10F) S10F {
	x.C = "not hinata"
	x.D = "not rocks"
	x.E = "not boruto"
	x.H = "not hinata"
	x.I = "not rocks"
	x.J = "not boruto"
	return x
}

// S10fByPointer is method to test passing pointers
func S10fByPointer() {
	x := S10F{A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto", F: "naruto", G: "rocks", H: "hinata", I: "rocks", J: "boruto"}
	s10fByPointer2(&x)
}

func s10fByPointer2(x *S10F) *S10F {
	x.A = "not naruto"
	x.B = "not rocks"
	x.F = "not naruto"
	x.G = "not rocks"
	return s10fByPointer3(x)
}

func s10fByPointer3(x *S10F) *S10F {
	x.C = "not hinata"
	x.D = "not rocks"
	x.E = "not boruto"
	x.H = "not hinata"
	x.I = "not rocks"
	x.J = "not boruto"
	return x
}
