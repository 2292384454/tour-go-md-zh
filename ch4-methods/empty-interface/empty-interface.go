package main

import "fmt"

func main() {
	// 指定了零个方法的接口值被称为空接口
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

// 空接口被用来处理未知类型的值
func describe(i interface{}) {
	fmt.Printf("(%v,%T)\n", i, i)
}
