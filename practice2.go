package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	x := uint8(dx)
	y := uint8(dy)

	z := (x + y) / 2
}

func main() {
	pic.Show(Pic)
}
