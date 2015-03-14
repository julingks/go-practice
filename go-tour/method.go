package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (c *Vertex) Abs() float64 {
	return math.Sqrt(c.X*c.X + c.Y*c.Y)
}

func (t *Vertex) Print() {
	fmt.Println("test", t.X)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Print()
}
