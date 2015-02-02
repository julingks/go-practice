package main

import (
	//"encoding/json"
	"fmt"
)

func main() {
	var i interface{}
	i = "a string"
	i = 2011
	i = 2.777

	r := i.(float64)
	fmt.Println(r)
	switch v := i.(type) {
	case int:
		fmt.Println("twice i is", v*2)
	case float64:
		fmt.Println("the reciprocal of i is", 1/v)
	case string:
		h := len(v) / 2
		fmt.Println("i swapped by halves is", v[h:]+v[:h])
	default:
		// i isn't one of the types above
	}

}
