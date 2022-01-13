package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	i := 0
	pre, beforePre := 0, 0
	return func() int {
		var cur int
		if i < 2 {
			cur = i
		} else {
			cur = pre + beforePre
		}
		beforePre, pre = pre, cur
		i++
		return cur
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
