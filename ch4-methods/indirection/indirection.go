package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// Scale 接收者为 *Vertex 类型的方法
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ScaleFunc 入参为 *Vertex 和 float64 类型的函数
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// kevinhwang 带指针参数的函数必须接受一个指针
// 而以指针为接收者的方法被调用时，接收者既能为值又能为指针
func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
