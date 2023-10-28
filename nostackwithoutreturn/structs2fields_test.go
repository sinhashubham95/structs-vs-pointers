package nostackwithoutreturn

import "testing"

func BenchmarkS2fByValue(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		S2fByValue()
	}
}

func BenchmarkS2fByPointer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		S2fByPointer()
	}
}
