package main

import "fmt"

func main() {
	z := []int{1, 2, 3}
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}
