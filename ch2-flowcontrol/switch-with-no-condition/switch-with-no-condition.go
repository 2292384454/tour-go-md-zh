package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { //kevinhwang: 没有条件的 switch 等价于 switch true
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
