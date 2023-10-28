package basic

// S5F is a 5 field struct.
type S5F struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
	E string `json:"e"`
}

// S5fByValue is method to test passing structs
func S5fByValue() {
	x := S5F{A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto"}
	x.A = "not naruto"
	x.B = "not rocks"
	x.C = "not hinata"
	x.D = "not rocks"
	x.E = "not boruto"
}

// S5fByPointer is method to test passing pointers
func S5fByPointer() {
	x := &S5F{A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto"}
	x.A = "not naruto"
	x.B = "not rocks"
	x.C = "not hinata"
	x.D = "not rocks"
	x.E = "not boruto"
}
