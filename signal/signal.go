package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time" // or "runtime"
)

func cleanup() {
	fmt.Println("cleanup!!!!!!!!!!!!!")
}

func main() {
	c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		sig := <-c
		switch sig {
		case syscall.SIGTERM:
			fmt.Println("SIGTERM")
		case os.Interrupt:
			fmt.Println("SIGINT")
		default:
			fmt.Println(sig)
		}
		cleanup()
		os.Exit(1)
	}()

	for {
		fmt.Println("sleeping...")
		time.Sleep(2 * time.Second)
	}
}
