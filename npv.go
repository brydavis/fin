package main

// TODO: May not work with NEGATIVE numbers

import "math"

func Round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	f = f * shift
	return math.Floor(f+.5) / shift
}

func NPV(value, discount float64, term int) float64 {
	return value / math.Pow((1.+discount), float64(term))
}

func PercentChange(former, current float64) float64 {
	return (current - former) / former
}
