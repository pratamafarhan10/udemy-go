package main

import "fmt"

func main() {
	c := make(chan<- int, 1)

	c <- 42

	fmt.Println(<-c) // this doesnt work because we can only send to the channel
}
