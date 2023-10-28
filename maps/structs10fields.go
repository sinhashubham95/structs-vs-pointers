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
	x := map[string]S10F{
		"a": {A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto", F: "naruto", G: "rocks", H: "hinata", I: "rocks", J: "boruto"},
		"b": {A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto", F: "naruto", G: "rocks", H: "hinata", I: "rocks", J: "boruto"},
	}
	s10fByValue2(x)
}

func s10fByValue2(x map[string]S10F) map[string]S10F {
	for k := range x {
		x[k] = S10F{
			A: "not naruto",
			B: "not rocks",
			C: x[k].C,
			D: x[k].D,
			E: x[k].E,
			F: "not naruto",
			G: "not rocks",
			H: x[k].H,
			I: x[k].I,
			J: x[k].J,
		}
	}
	return s10fByValue3(x)
}

func s10fByValue3(x map[string]S10F) map[string]S10F {
	for k := range x {
		x[k] = S10F{
			A: x[k].A,
			B: x[k].B,
			C: "not hinata",
			D: "not rocks",
			E: "not boruto",
			F: x[k].F,
			G: x[k].G,
			H: "not hinata",
			I: "not rocks",
			J: "not boruto",
		}
	}
	return x
}

// S10fByPointer is method to test passing pointers
func S10fByPointer() {
	x := map[string]*S10F{
		"a": {A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto", F: "naruto", G: "rocks", H: "hinata", I: "rocks", J: "boruto"},
		"b": {A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto", F: "naruto", G: "rocks", H: "hinata", I: "rocks", J: "boruto"},
	}
	s10fByPointer2(x)
}

func s10fByPointer2(x map[string]*S10F) map[string]*S10F {
	for _, v := range x {
		v.A = "not naruto"
		v.B = "not rocks"
		v.F = "not naruto"
		v.G = "not rocks"
	}
	return s10fByPointer3(x)
}

func s10fByPointer3(x map[string]*S10F) map[string]*S10F {
	for _, v := range x {
		v.C = "not hinata"
		v.D = "not rocks"
		v.E = "not boruto"
		v.H = "not hinata"
		v.I = "not rocks"
		v.J = "not boruto"
	}
	return x
}
