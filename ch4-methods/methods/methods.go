package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Abs 书写形式上与一般函数的区别：在方法名前面加上了接收者
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
