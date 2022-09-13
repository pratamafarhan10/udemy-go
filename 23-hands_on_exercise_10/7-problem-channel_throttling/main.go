package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	// send
	go populate(c)

	// receive
	receive(c)
	fmt.Println("About to exit")
}

func populate(c chan<- int) {
	const goroutine = 10
	var wg sync.WaitGroup
	wg.Add(goroutine)

	go func() {
		for i := 0; i < 100; i++ {
			c <- timesTen(i)
		}
		defer wg.Done()
		close(c)
	}()
	wg.Wait()
}

func receive(c <-chan int) {
	for val := range c {
		fmt.Println(val)
	}
}

func timesTen(num int) int {
	return num * 10
}
