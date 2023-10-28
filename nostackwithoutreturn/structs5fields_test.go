package nostackwithoutreturn

import "testing"

func BenchmarkS5fByValue(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		S5fByValue()
	}
}

func BenchmarkS5fByPointer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		S5fByPointer()
	}
}
