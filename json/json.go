package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}

	b, _ := json.Marshal(m)
	fmt.Println(string(b))

	var m2 Message

	json.Unmarshal(b, &m2)

	fmt.Println(m2)
}
