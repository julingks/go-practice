package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s)
	fmt.Println("len : ", len(s)) //5
	fmt.Println("cap : ", cap(s)) //5

	s = s[2:4]

	fmt.Println("s[2:4] : ", s) // 2 3

	fmt.Println("len : ", len(s)) //2
	fmt.Println("cap : ", cap(s)) //4
}
