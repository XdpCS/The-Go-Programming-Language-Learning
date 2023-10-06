package weightconv

func PToK(p Pound) Kilogram { return Kilogram(p * OnePoundInKilograms) }
func KToP(k Kilogram) Pound { return Pound(k / OnePoundInKilograms) }
