package main

import "fmt"

func main() {
	// kevinhwang: 函数外的每个语句都必须以关键字开始（var, func 等等）
	// := 结构不能在函数外使用
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
