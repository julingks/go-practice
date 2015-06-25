package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

const (
	timeout = time.Second * 10
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	ctx = context.WithValue(ctx, "key", "value")

	cancel()
	select {
	case <-time.After(20 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
	fmt.Println(ctx.Value("key").(string))
}
