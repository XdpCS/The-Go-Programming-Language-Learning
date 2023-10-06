package main

import (
	"fmt"
	"github.com/XdpCs/ch2/tempconv"
)

func main() {
	c := tempconv.Celsius(1118.0)
	k := tempconv.CtoK(c)
	fmt.Println(k)
	c = tempconv.KtoC(k)
	fmt.Println(c)
	f := tempconv.Fahrenheit(1118.0)
	k = tempconv.FtoK(f)
	fmt.Println(k)
	f = tempconv.KtoF(k)
	fmt.Println(f)
}
