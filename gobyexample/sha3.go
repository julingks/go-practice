package main

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

func main() {
	hash := sha3.New256()
	hash.Write([]byte("white house"))
	bs := hash.Sum([]byte{})
	fmt.Printf("%v", bs)
}
