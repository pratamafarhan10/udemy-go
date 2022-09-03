package main

import "fmt"

func main() {
	x := make(chan int)

	go func() {
		x <- 42 // this will work because the channel run in another go routine
	}()

	fmt.Println(<-x)
}
