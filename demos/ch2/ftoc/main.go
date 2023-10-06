package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // 32°F = 0°C
	// %g 根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF)) // 212°F = 100°C

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
