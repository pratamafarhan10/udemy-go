package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int)
	cs := make(chan<- int)

	fmt.Printf("Basic channel: %T\n", c)
	fmt.Printf("Can only recive from channel: %T\n", cr)
	fmt.Printf("Can only send to channel: %T\n", cs)
}
