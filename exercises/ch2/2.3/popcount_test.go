package popcount

import "testing"

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(1118)
	}
}

func BenchmarkPopCountOneExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountOneExpression(1118)
	}
}
