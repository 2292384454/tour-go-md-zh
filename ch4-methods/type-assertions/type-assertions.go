package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// 该语句断言接口值 i 保存了具体类型 string ，并将其底层类型为 string 的值赋予变量 s
	// 如果 i 未保存 T 类型的变量值，该语句就会触发一个 panic
	s := i.(string)
	fmt.Printf("(%v,%T)\n", s, s)

	// 该语句判断接口值 i 是否保存了 string 类型，将判断结果返回到第二个返回值 ok 中
	// 若 i 保存了一个 string ，那么 s 将会是其底层值，而 ok 为 true
	// 否则， ok 将为 false 而 s 将为 string 类型的零值，即空字符串
	s, ok := i.(string)
	fmt.Printf("(%v,%T,%v)\n", s, s, ok)

	f, ok := i.(float64)
	fmt.Printf("(%v,%T,%v)\n", f, f, ok)

	f = i.(float64) // 报错(panic)
	fmt.Printf("(%v,%T)\n", f, f)
}
