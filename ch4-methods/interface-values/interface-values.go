package main

import (
	"fmt"
	"math"
)

// I 定义接口I， I 有方法 M()
type I interface {
	M()
}

// T 定义结构体 T ，包含一个变量 S
type T struct {
	S string
}

// M 为结构体 T 实现方法 M()
func (t *T) M() {
	fmt.Println(t.S)
}

// F 定义类型F
type F float64

// M 为结构体 F 实现方法 M()
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
