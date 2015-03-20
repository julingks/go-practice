package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.SplitAfter("1111,222,3,44,55,6666,7777777,88,9,10", ","))
	fmt.Println(strings.Split("1111,222,3,44,55,6666,7777777,88,9,10", ","))

}
