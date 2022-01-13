package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")
	// KevinHwang： Go 中格式化输出时间的方法 Format("2006/01/02 15:04:05") ，此处时间为固定常量
	fmt.Println("The time is", time.Now().Format("2006/01/02 15:04:05"))
}
