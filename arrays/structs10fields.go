package arrays

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
	x := []S10F{{A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto", F: "naruto", G: "rocks", H: "hinata", I: "rocks", J: "boruto"}, {A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto", F: "naruto", G: "rocks", H: "hinata", I: "rocks", J: "boruto"}}
	s10fByValue2(x)
}

func s10fByValue2(x []S10F) []S10F {
	x[0].A = "not naruto"
	x[0].B = "not rocks"
	x[0].F = "not naruto"
	x[0].G = "not rocks"
	x[1].A = "not naruto"
	x[1].B = "not rocks"
	x[1].F = "not naruto"
	x[1].G = "not rocks"
	return s10fByValue3(x)
}

func s10fByValue3(x []S10F) []S10F {
	x[0].C = "not hinata"
	x[0].D = "not rocks"
	x[0].E = "not boruto"
	x[0].H = "not hinata"
	x[0].I = "not rocks"
	x[0].J = "not boruto"
	x[1].C = "not hinata"
	x[1].D = "not rocks"
	x[1].E = "not boruto"
	x[1].H = "not hinata"
	x[1].I = "not rocks"
	x[1].J = "not boruto"
	return x
}

// S10fByPointer is method to test passing pointers
func S10fByPointer() {
	x := []*S10F{{A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto", F: "naruto", G: "rocks", H: "hinata", I: "rocks", J: "boruto"}, {A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto", F: "naruto", G: "rocks", H: "hinata", I: "rocks", J: "boruto"}}
	s10fByPointer2(x)
}

func s10fByPointer2(x []*S10F) []*S10F {
	x[0].A = "not naruto"
	x[0].B = "not rocks"
	x[0].F = "not naruto"
	x[0].G = "not rocks"
	x[1].A = "not naruto"
	x[1].B = "not rocks"
	x[1].F = "not naruto"
	x[1].G = "not rocks"
	return s10fByPointer3(x)
}

func s10fByPointer3(x []*S10F) []*S10F {
	x[0].C = "not hinata"
	x[0].D = "not rocks"
	x[0].E = "not boruto"
	x[0].H = "not hinata"
	x[0].I = "not rocks"
	x[0].J = "not boruto"
	x[1].C = "not hinata"
	x[1].D = "not rocks"
	x[1].E = "not boruto"
	x[1].H = "not hinata"
	x[1].I = "not rocks"
	x[1].J = "not boruto"
	return x
}
