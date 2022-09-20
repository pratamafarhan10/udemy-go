package main

import "fmt"

func main() {
	x := make(chan int)

	// send
	v := send(x)

	// receive
	receive(v)
}

// send receive a bi-directional channel as a parameter,
// send value to the channel,
// and returns a receive-only channel
func send(c chan int) <-chan int {
	go func() {
		c <- 42
	}()
	return c
}

// receive receive a receive-only channel
// and prints the value pulled from the channel
func receive(c <-chan int) {
	fmt.Println(<-c)
}
