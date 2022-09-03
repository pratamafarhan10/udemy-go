package main

import "fmt"

func main() {
	x := make(chan int, 2)

	x <- 42
	x <- 43

	fmt.Println(<-x)
	fmt.Println(<-x)

	x <- 9
	x <- 10

	fmt.Println(<-x)
	fmt.Println(<-x)
}
