package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/XdpCs/ch2/lenconv"
	"github.com/XdpCs/ch2/tempconv"
	"github.com/XdpCs/ch2/weightconv"
)

func main() {
	numbers := os.Args[1:]
	number := 0.0
	if len(numbers) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			f, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "convert: %v\n", err)
				os.Exit(1)
			}
			number = f
			break
		}
	} else {
		f, err := strconv.ParseFloat(numbers[0], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "convert: %v\n", err)
			os.Exit(1)
		}
		number = f
	}

	fmt.Printf("%s = %s, %s = %s\n",
		tempconv.Fahrenheit(number), tempconv.FToC(tempconv.Fahrenheit(number)),
		tempconv.Celsius(number), tempconv.CToF(tempconv.Celsius(number)))
	fmt.Printf("%s = %s, %s = %s\n",
		weightconv.Pound(number), weightconv.PToK(weightconv.Pound(number)),
		weightconv.Kilogram(number), weightconv.KToP(weightconv.Kilogram(number)))
	fmt.Printf("%s = %s, %s = %s\n",
		lenconv.Inch(number), lenconv.InToM(lenconv.Inch(number)),
		lenconv.Meter(number), lenconv.MtoIn(lenconv.Meter(number)))
}
