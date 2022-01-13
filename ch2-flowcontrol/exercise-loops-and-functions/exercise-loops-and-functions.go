package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	fmt.Println(z)
	for step := (z*z - x) / (2 * z); math.Abs(step) > 1e-10; step = (z*z - x) / (2 * z) {
		z -= step
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
