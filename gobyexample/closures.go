package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1 // 익명 함수 안에서 변수 정의없이 사용할 수 있다.
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()

	fmt.Println(newInts())
}
