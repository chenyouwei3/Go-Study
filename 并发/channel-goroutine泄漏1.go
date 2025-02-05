package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	var ch chan struct{}
	go func() {
		ch <- struct{}{}
	}()

	time.Sleep(time.Second)
}
