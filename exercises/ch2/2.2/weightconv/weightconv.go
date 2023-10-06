package weightconv

import "fmt"

type Pound float64
type Kilogram float64

const OnePoundInKilograms = 0.45359237

func (p Pound) String() string    { return fmt.Sprintf("%g kg", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g lb", k) }
