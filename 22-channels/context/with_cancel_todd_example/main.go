package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Num of error:", ctx.Err())
	fmt.Println("Num of goroutine:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)

	fmt.Println("Num of error:", ctx.Err())
	fmt.Println("Num of goroutine:", runtime.NumGoroutine())

	cancel()

	time.Sleep(time.Second * 2)

	fmt.Println("Num of error:", ctx.Err())
	fmt.Println("Num of goroutine:", runtime.NumGoroutine())
}
