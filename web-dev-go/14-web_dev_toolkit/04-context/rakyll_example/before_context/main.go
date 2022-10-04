package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("NUM OF GOROUTINE BEFORE BREAK:", runtime.NumGoroutine())

	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	time.Sleep(10 * time.Second)
	fmt.Println("NUM OF GOROUTINE AFTER BREAK:", runtime.NumGoroutine())
}

func gen() <-chan int {
	ch := make(chan int)

	go func() {
		var n int
		for {
			ch <- n
			n++
		}
	}()

	return ch
}
