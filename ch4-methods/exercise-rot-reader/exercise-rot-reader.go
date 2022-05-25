package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// kevinhwang 通过实现 Read() 方法让 rot13Reader 实现 io.Reader 接口
func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i := 0; i <= n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

// rot13 rot13加密算法
func rot13(x byte) byte {
	if x < 'A' || (x > 'Z' && x < 'a') || x > 'z' {
		return x
	}
	if x <= 'Z' {
		return (x-'A'+13)%('Z'-'A'+1) + 'A'
	} else {
		return (x-'a'+13)%('z'-'a'+1) + 'a'
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
