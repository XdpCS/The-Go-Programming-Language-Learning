package lenconv

func MtoIn(m Meter) Inch { return Inch(m / OneInchPerMeter) }
func InToM(i Inch) Meter { return Meter(i * OneInchPerMeter) }
