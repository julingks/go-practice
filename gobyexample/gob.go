package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

// 네트웍을 통해서 데이터 구조를 전송하거나 파일에 저장하기 위해서,
// 인코딩 또는 디코딩을 해야 한다.
// JSON, XML, protocol buffers 등이 있지만, go는 gob 패키지를 제공한다.

func main() {
	// 인코더와 디코더를 초기화한다.
	// 보통 enc와 dec는 네크웍 연결에 속박된다.
	// encoder와 decoder는 다른 프로세스에서 실행된다.

	var network bytes.Buffer
	// NewEncoder와 NewDecoder는  io.Writer를 받는다.
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	// 값을 인코딩한다.
	p := P{3, 4, 5, "Pythagoras"}
	err := enc.Encode(p)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// 값을 디코딩한다.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Println("send : ", p)
	fmt.Printf("receive : %q: {%d,%d}\n", q.Name, *q.X, *q.Y)
}
