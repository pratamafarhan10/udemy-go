package main

import "fmt"

func main() {
	x := make(chan int, 1)

	x <- 42
	x <- 43 // This wouldnt work because channel only have 1 room to store a value

	fmt.Println(<-x)
}
