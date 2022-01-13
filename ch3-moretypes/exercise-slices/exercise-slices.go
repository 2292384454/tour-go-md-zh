package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/tour/pic"
	"os"
)

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for x := range image {
		image[x] = make([]uint8, dx)
		for y := range image[x] {
			image[x][y] = uint8(x % (y + 1))
		}
	}
	return image
}

func main() {
	str := pic.Show(Pic)
	// base64 字符串解码成图片
	dist, _ := base64.StdEncoding.DecodeString(str)
	// 写入文件
	f, err := os.OpenFile("ch3-moretypes/exercise-slices-output.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	if err == nil {
		f.Write(dist)
		fmt.Println("图片已保存至 ch3-moretypes/exercise-slices-output.png")
	}
}
