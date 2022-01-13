package main

import (
	"fmt"
	"math"
)

// Sqrt kevinhwang: 采用牛顿迭代法计算平方根，当迭代梯度小于 1e-10 时返回结果
func Sqrt(x float64) float64 {
	z := x / 2
	//fmt.Println(z)
	for step := (z*z - x) / (2 * z); math.Abs(step) > 1e-10; step = (z*z - x) / (2 * z) {
		z -= step
		//fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println("Sqrt(2)=", Sqrt(2))

	// kevinhwang 与math.Sqrt()对比
	fmt.Println("math.Sqrt(2)=", math.Sqrt(2))
	fmt.Println("Sqrt(2)-math.Sqrt(2)=", Sqrt(2)-math.Sqrt(2))
}
