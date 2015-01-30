package main

import "fmt"

func main() {

	a := []string{"John", "Paul"}
	b := []string{"George", "Ringo", "Pete"}
	a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
	// a == []string{"John", "Paul", "George", "Ringo", "Pete"}

	fmt.Println(a)

	c := make([]int, 10)
	for k, v := range c {
		fmt.Print(k, v, "   ")
	}
	fmt.Println(cap(c))
	d := make([]int, 10, 20)
	for k, v := range d {
		fmt.Print(k, v, "   ")
	}
	fmt.Println(cap(d))
}
