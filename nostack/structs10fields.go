package basic

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
	_ = S10F{A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto", F: "naruto", G: "rocks", H: "hinata", I: "rocks", J: "boruto"}
}

// S10fByPointer is method to test passing pointers
func S10fByPointer() {
	_ = &S10F{A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto", F: "naruto", G: "rocks", H: "hinata", I: "rocks", J: "boruto"}
}
