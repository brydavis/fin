package main

func Simple(principal, rate float64, term int) float64 {
	return principal + (principal * rate * float64(term))
}

func Compound(principal, rate float64, term int) float64 {
	if term == 0 {
		return principal
	} else {
		return Compound(principal+(principal*rate), rate, term-1)
	}
}
