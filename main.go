package main

import (
	"fmt"

	// "github.com/brydavis/fin"
)

func main() {
	p := 1000.0 // principle
	r := 0.05   // interest rate
	y := 35     // years
	v := Compound(p, r, y)

	fmt.Printf("original principle: $%0.2f\tbalance @ %d years: $%0.2f\n", p, y, v) // 5516.015368

	fmt.Println(NPV(300000., 0.06, 1))
	fmt.Println(PercentChange(500, 600))

}
