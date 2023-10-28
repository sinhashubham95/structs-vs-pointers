package arrays

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
	x := map[string]S5F{
		"a": {A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto"},
		"b": {A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto"},
	}
	s5fByValue2(x)
}

func s5fByValue2(x map[string]S5F) map[string]S5F {
	for k := range x {
		x[k] = S5F{
			A: "not naruto",
			B: "not rocks",
			C: x[k].C,
			D: x[k].D,
			E: x[k].E,
		}
	}
	return s5fByValue3(x)
}

func s5fByValue3(x map[string]S5F) map[string]S5F {
	for k := range x {
		x[k] = S5F{
			A: x[k].A,
			B: x[k].B,
			C: "not hinata",
			D: "not rocks",
			E: "not boruto",
		}
	}
	return x
}

// S5fByPointer is method to test passing pointers
func S5fByPointer() {
	x := map[string]*S5F{
		"a": {A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto"},
		"b": {A: "naruto", B: "rocks", C: "hinata", D: "rocks", E: "boruto"},
	}
	s5fByPointer2(x)
}

func s5fByPointer2(x map[string]*S5F) map[string]*S5F {
	for k := range x {
		x[k].A = "not naruto"
		x[k].B = "not rocks"
	}
	return s5fByPointer3(x)
}

func s5fByPointer3(x map[string]*S5F) map[string]*S5F {
	for k := range x {
		x[k].C = "not hinata"
		x[k].D = "not rocks"
		x[k].E = "not boruto"
	}
	return x
}
