package main

import "fmt"

func main() {
	p := 1000.0
	r := 0.12
	t := 100
	v := Compound(p, r, t)

	// fmt.Printf("%f\n", Round(v, 3))
	// fmt.Println(v)
	fmt.Printf("%f\n", v)

}
