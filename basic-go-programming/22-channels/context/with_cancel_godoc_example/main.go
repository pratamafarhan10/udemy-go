package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	for val := range gen(ctx) {
		fmt.Println(val)
		if val == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	n := 1
	dst := make(chan int)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()

	return dst
}
