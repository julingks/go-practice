package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
)

type Image struct{}

func (f *Image) ColorModel() color.Model {
	return nil
}

func (f *Image) Bounds() image.Rectangle {
	return nil
}

func (f *Image) At(x, y int) color.Color {
	return nil
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
