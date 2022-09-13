package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go populate(c)

	// receive
	receive(c)
	fmt.Println("About to exit")
}

func populate(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	for val := range c {
		fmt.Println(val)
	}
}
