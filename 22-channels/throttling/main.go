package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	go populate(c)

	for v := range c {
		fmt.Println(v)
	}
}

// populate populates the channel
func populate(c chan int) {
	var wg sync.WaitGroup

	// goroutines represents the number of goroutine can run
	const goroutines = 10

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(n int) {
			c <- n
			defer wg.Done()
		}(i)
	}

	wg.Wait()

	close(c)
}
