package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	"os"
)

type Image struct {
	rect image.Rectangle
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return i.rect
}
func (i Image) At(x, y int) color.Color {
	v := uint8(x * y)
	return color.RGBA{R: v, G: v, B: 255, A: 255}
}

func main() {
	m := Image{
		rect: image.Rectangle{
			Min: image.Point{},
			Max: image.Point{X: 256, Y: 256}}}
	str := pic.ShowImage(m)

	// kevinhwang 原本的 pic.ShowImage() 行数将返回一个 base64 字符串，由于无法在控制台显示图片，所以对该函数做了修改，将 base64 字符串返回后转成图片再保存到本地

	// base64 字符串解码成图片
	dist, _ := base64.StdEncoding.DecodeString(str)
	// 写入文件
	f, err := os.OpenFile("ch4-methods/exercise-images/exercise-images-output.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	if err == nil {
		f.Write(dist)
		fmt.Println("图片已保存至 ch4-methods/exercise-images/exercise-images-output.png")
	}
}
