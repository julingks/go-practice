package main

import "fmt"

func main() {
	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:]
	fmt.Printf("e : %s\n", e)

	e[1] = 'm'

	fmt.Printf("d : %s\n", d)

	a := []string{"1", "2", "3"}
	b := a[:2]
	b[1] = "4"
	fmt.Println(a)
	fmt.Println(b)

	s := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(len(s))
	s = s[:cap(s)]
	fmt.Println(s)
	fmt.Println(len(s))
}
