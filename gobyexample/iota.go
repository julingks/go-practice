package main

import "fmt"

type ByteSize float64

/*
	iota는상수의enumerator 이다. 0 부터 시작한다.
*/

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	A = iota
	B
	C
	D
	E
)

func main() {
	fmt.Printf("kb %f\n", KB)
	fmt.Printf("mb %f\n", MB)
	fmt.Printf("gb %f\n", GB)
	fmt.Printf("tb %f\n", TB)
	fmt.Printf("pb %f\n", PB)
	fmt.Printf("eb %f\n", EB)
	fmt.Printf("zb %f\n", ZB)
	fmt.Printf("yb %f\n", YB)

	fmt.Println("A : ", A)
	fmt.Println("B : ", B)
	fmt.Println("C : ", C)
	fmt.Println("D : ", D)
	fmt.Println("E : ", E)
}
