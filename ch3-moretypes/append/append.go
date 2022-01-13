package main

import "fmt"

// kevinhwang
// 调用append 函数时 capacity如何增长：
// 1.先将旧的slice容量乘以2，如果乘以2后的容量仍小于新的slice容量，则取新的slice容量(append多个elems)
// 2.如果新slice小于等于旧slice容量的2倍，则取旧slice容量乘以2
// 3.如果旧的slice容量大于1024，则新slice容量取旧slice容量乘以对应的倍数(不小于1.25)
// 代码后面还会对 newcap 进行roundup，比如在64位平台，newcap 是奇数的话就会+1

func main() {
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
