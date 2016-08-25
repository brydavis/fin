package main

import "math"

func Simple(present, rate float64, time int) float64 {
	// V = P + (P * R * T)
	return present + (present * rate * float64(time))
}

func Compound(present, rate float64, time int) float64 {
	if time == 0 {
		return present
	} else {
		return Compound(present+(present*rate), rate, time-1)
	}
}

// TODO: May not work with NEGATIVE numbers
func Round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	f = f * shift
	return math.Floor(f+.5) / shift
}
