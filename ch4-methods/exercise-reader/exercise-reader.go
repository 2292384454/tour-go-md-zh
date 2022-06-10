package main

import (
	"fmt"
	"golang.org/x/tour/reader"
)

// MyReader 产生一个 ASCII 字符 'A' 的无限流。
type MyReader struct{}

// MyReaderError kevinhwang 作为切片容量不足的 error
type MyReaderError int

func (e MyReaderError) Error() string {
	return fmt.Sprintf("Array with length %d is too short", e)
}

// Read 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func (MyReader) Read(arr []byte) (int, error) {
	l := len(arr)
	if l < 1 {
		return l, MyReaderError(l)
	}
	for i := range arr {
		arr[i] = byte('A')
	}
	return l, nil
}

func main() {
	reader.Validate(MyReader{})

	// kevinhwang:test
	myReader := MyReader{}
	arr := make([]byte, 8)
	n, err := myReader.Read(arr)
	fmt.Println(n, err)
}
