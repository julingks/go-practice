package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

func server() {
	// 포트 대기
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("before accept")

	for {
		// 연결 수락
		c, err := ln.Accept()
		fmt.Println("accepted")
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handlerServerConnection(c)
	}
}

func handlerServerConnection(c net.Conn) {
	// 메시지 수신
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}

	c.Close()
}

func client() {
	time.Sleep(time.Second * 3)
	// 서버에 연결
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 메시지 송신
	msg := "Hello World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	go server()
	go client()
	var input string
	fmt.Scanln(&input)
}
