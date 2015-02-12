package main

import (
	"fmt"
	"time"
)

import "sync/atomic"

func main() {
	var ops uint64 = 0

	for i := 0; i < 5; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				//runtime.Gosched()
			}
		}()
		fmt.Println(i)
	}

	fmt.Println("A")
	time.Sleep(time.Second)
	fmt.Println("B")

	opsFinal := atomic.LoadUint64(&ops)

	fmt.Println("ops:", opsFinal)
}
