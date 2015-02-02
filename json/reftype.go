package main

import (
	"encoding/json"
	"fmt"
)

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

type Foo struct {
	Bar *bar
}

type IncomingMessage struct {
	Cmd *Command
	Msg *Message
}

func main() {

	var m FamilyMember
	err := json.Unmarshal(b, &m)
}
