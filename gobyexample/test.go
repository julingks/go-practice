package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s[1:2]) //1

	fmt.Println(s[1:5]) //1 2 3 4

	fmt.Println(s[1:6]) //1 2 3 4 5
	fmt.Println(s[0:6]) //0 1 2 3 4 5

	fmt.Println(s[1:3]) //1 2

	fmt.Println(s[1]) //1
	fmt.Println(s[3]) //3
}
