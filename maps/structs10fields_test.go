package arrays

import "testing"

func BenchmarkS10fByValue(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		S10fByValue()
	}
}

func BenchmarkS10fByPointer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		S10fByPointer()
	}
}
