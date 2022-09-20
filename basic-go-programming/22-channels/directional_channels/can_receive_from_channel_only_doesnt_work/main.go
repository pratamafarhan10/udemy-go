package main

import "fmt"

func main() {
	c := make(<-chan int, 1)

	c <- 42 // this doesnt work because we can only receive from channel

	fmt.Println(<-c)
}
