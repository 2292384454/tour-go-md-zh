package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// kevinhwang: 用当前时间戳作为随机数种子，实现真正地随机数生成
	// rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))
}
