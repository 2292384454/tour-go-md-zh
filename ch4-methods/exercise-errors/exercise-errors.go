package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot`Sqrt`negative`number:%g\n", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return -1, ErrNegativeSqrt(x)
	}
	z := x / 2
	for step := (z*z - x) / (2 * z); math.Abs(step) > 1e-10; step = (z*z - x) / (2 * z) {
		z -= step
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
