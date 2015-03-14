package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for i := range make([]int, 10) {
		fmt.Println(i)
	}

	m := make(map[int]string)

	m[1] = "value1"
	m[2] = "value2"

	for k, v := range m {
		fmt.Println(k, v)
	}
}
