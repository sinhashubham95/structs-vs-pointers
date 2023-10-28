package basic

// S2F is a 2 field struct.
type S2F struct {
	A string `json:"a"`
	B string `json:"b"`
}

// S2fByValue is method to test passing structs
func S2fByValue() {
	x := S2F{A: "naruto", B: "rocks"}
	x.A = "not naruto"
	x.B = "not rocks"
}

// S2fByPointer is method to test passing pointers
func S2fByPointer() {
	x := &S2F{A: "naruto", B: "rocks"}
	x.A = "not naruto"
	x.B = "not rocks"
}
