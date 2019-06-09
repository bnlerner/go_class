package main

import (
	"fmt"
	"math"
)

func main() {
	alpha := GenDisplaceFn(10, 2, 1)
	alpha(1)
	alpha(2)
	alpha(3)

	bravo := GenDisplaceFn(20, 0, 0)
	bravo(1)
	bravo(2)
	bravo(3)
}

/*
* Calculates displacement using the formula
* x = (0.5 * a * t^2) + (vi * t) + xi
 */
func GenDisplaceFn(a float64, vi float64, xi float64) func(t float64) {
	return func(t float64) {
		p1 := 0.5 * a * math.Pow(t, 2)
		p2 := vi * t
		p3 := xi

		x := p1 + p2 + p3
		fmt.Printf("At t = %.2f, displacement = %.5f\n", t, x)
	}
}
