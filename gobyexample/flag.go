package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	maxp := flag.Int("max", 6, "최댓값")

	flag.Parse()

	fmt.Println("maxp : ", *maxp)
	fmt.Println(rand.Intn(*maxp))
}
