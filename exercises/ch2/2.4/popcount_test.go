package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(1118)
	}
}

func BenchmarkPopCountOffset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountOffset(1118)
	}
}
