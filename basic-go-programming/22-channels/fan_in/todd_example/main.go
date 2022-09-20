package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	// send
	go send(even, odd)

	// receive
	go receive(even, odd, fanin)

	for val := range fanin {
		fmt.Println(val)
	}

	fmt.Println("======= About to exit")
}

func send(even, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	close(even)
	close(odd)
}

func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for val := range even {
			fanin <- val
		}
		defer wg.Done()
	}()

	go func() {
		for val := range odd {
			fanin <- val
		}
		defer wg.Done()
	}()

	wg.Wait()

	close(fanin)
}
