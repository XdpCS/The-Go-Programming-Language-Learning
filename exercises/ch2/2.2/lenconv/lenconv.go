package lenconv

import "fmt"

type Inch float64
type Meter float64

const OneInchPerMeter = 0.0254

func (i Inch) String() string  { return fmt.Sprintf("%g in", i) }
func (m Meter) String() string { return fmt.Sprintf("%g m", m) }
