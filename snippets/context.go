package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

const (
	timeout = time.Millisecond * 100
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
	defer cancel()
}
