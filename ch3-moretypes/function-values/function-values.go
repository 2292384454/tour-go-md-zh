package main

import (
	"fmt"
	"math"
)

// kevinhwang： compute 函数将使用传入的 fn 函数计算 x 和 y 的运算结果
// 例如：compute(math.Pow, 2, 10) 将返回 math.Pow(2, 10)
func compute(fn func(float64, float64) float64, x float64, y float64) float64 {
	return fn(x, y)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot, 5, 12))
	fmt.Println(compute(math.Pow, 2, 10))
}
