package main

import "fmt"

func main() {
	x := make(chan int)

	x <- 42 // this alone wouldn't work because channel blocks the main go routine

	fmt.Println(<-x)
}
