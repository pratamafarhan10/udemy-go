package main

import (
	"fmt"
)

func main() {
	// solution 1
	// c := make(chan int, 1)

	// solution 2
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
