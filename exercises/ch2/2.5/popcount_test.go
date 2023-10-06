package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(1118)
	}
}

func BenchmarkPopCountXAndXSubOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountXAndXSubOne(1118)
	}
}
