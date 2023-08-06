package echo

import (
	"os"
	"strings"
)

func Echo0() string {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	return s
}

func Echo1() string {
	return strings.Join(os.Args, " ")
}
