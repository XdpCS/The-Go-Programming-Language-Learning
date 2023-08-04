package echo_test

import (
	"os"
	"strings"
	"testing"
)

func BenchmarkEcho0(b *testing.B) {
	os.Args = []string{"fyy", "xdp"}
	for i := 0; i < b.N; i++ {
		if echo0() != "fyy xdp" {
			b.Error("echo0() != \"fyy xdp\"")
		}
	}
}

func BenchmarkEcho1(b *testing.B) {
	os.Args = []string{"fyy", "xdp"}
	for i := 0; i < b.N; i++ {
		if echo1() != "fyy xdp" {
			b.Error("echo1() != \"fyy xdp\"")
		}
	}
}

func echo0() string {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo1() string {
	return strings.Join(os.Args, " ")
}
