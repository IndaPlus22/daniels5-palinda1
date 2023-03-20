package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	prev := 0.0
	for math.Abs(prev-z) > 0.0000001 {
		prev = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)

	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
