package echo

import (
	"os"
	"testing"
)

var testCase []string = []string{"fyy", "xdp"}

func BenchmarkEcho0(b *testing.B) {
	os.Args = testCase
	for i := 0; i < b.N; i++ {
		if Echo0() != "fyy xdp" {
			b.Error("echo0() != \"fyy xdp\"")
		}
	}
}

func BenchmarkEcho1(b *testing.B) {
	os.Args = testCase
	for i := 0; i < b.N; i++ {
		if Echo1() != "fyy xdp" {
			b.Error("echo1() != \"fyy xdp\"")
		}
	}
}
