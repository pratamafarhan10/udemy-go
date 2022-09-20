package main

import (
	"fmt"
	"sync"
)

func write(channel chan<- string) {
	defer wg.Done()
	channel <- "hello there"
}

func read(channel <-chan string) {
	defer wg.Done()
	fmt.Printf(<-channel)
}

var wg sync.WaitGroup

func main() {
	c := make(chan string)

	wg.Add(2)

	go write(c)
	go read(c)

	wg.Wait()
}
