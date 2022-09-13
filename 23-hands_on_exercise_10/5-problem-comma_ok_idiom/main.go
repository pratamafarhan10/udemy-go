package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 1
	}()

	v, open := <-c
	fmt.Println(v, open)

	close(c)

	v, open = <-c
	fmt.Println(v, open)
}
