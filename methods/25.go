package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	rect := image.Rect(0, 0, 255, 255)
	return rect
}

func (img Image) At(x, y int) color.Color {
	v := func1(x, y)
	return color.RGBA{v, v, 255, 255}
}

func func1(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func func2(x, y int) uint8 {
	return uint8(x * y)
}

func func3(x, y int) uint8 {
	return uint8(x ^ y)
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
