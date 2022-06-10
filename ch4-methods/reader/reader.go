package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// 实例化 Reader 对象
	r := strings.NewReader("Hello, Reader!")
	// 创建接收字节流的字节数组
	b := make([]byte, 8)
	// 在死循环中读 Reader
	for {
		// 返回值 n 表示读了多少个字节， err 表示读的过程中发生的错误
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		// 如果已经读到文件尾，就会触发 io.EOF 错误
		if err == io.EOF {
			break
		}
	}
}
