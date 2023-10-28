package arrays

// S2F is a 2 field struct.
type S2F struct {
	A string `json:"a"`
	B string `json:"b"`
}

// S2fByValue is method to test passing structs
func S2fByValue() {
	x := map[string]S2F{
		"a": {A: "naruto", B: "rocks"},
		"b": {A: "naruto", B: "rocks"},
	}
	s2fByValue2(x)
}

func s2fByValue2(x map[string]S2F) map[string]S2F {
	for k := range x {
		x[k] = S2F{
			A: "not naruto",
			B: x[k].B,
		}
	}
	return s2fByValue3(x)
}

func s2fByValue3(x map[string]S2F) map[string]S2F {
	for k := range x {
		x[k] = S2F{
			A: x[k].A,
			B: "not rocks",
		}
	}
	return x
}

// S2fByPointer is method to test passing pointers
func S2fByPointer() {
	x := map[string]*S2F{
		"a": {A: "naruto", B: "rocks"},
		"b": {A: "naruto", B: "rocks"},
	}
	s2fByPointer2(x)
}

func s2fByPointer2(x map[string]*S2F) map[string]*S2F {
	for k := range x {
		x[k].A = "not naruto"
	}
	return s2fByPointer3(x)
}

func s2fByPointer3(x map[string]*S2F) map[string]*S2F {
	for k := range x {
		x[k].B = "not rocks"
	}
	return x
}
