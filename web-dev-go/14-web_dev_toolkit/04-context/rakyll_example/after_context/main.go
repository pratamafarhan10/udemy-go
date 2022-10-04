package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("NUM OF GOROUTINE BEFORE BREAK:", runtime.NumGoroutine())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}
	time.Sleep(10 * time.Second)
	fmt.Println("NUM OF GOROUTINE AFTER BREAK:", runtime.NumGoroutine())
}

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- n
				n++
			}
		}
	}()

	return ch
}
