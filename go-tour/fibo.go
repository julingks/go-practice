package main

import "fmt"

func fibonacci() func() int {
	x1 := 0
	x2 := 1
	return func() int {
		ret := x1 + x2
		x1 = x2
		x2 = ret
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
