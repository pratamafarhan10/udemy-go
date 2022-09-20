package main

import "fmt"

func main() {
	x := make(chan int, 1) // allow 1 value to sit in there

	x <- 42 // got buffer of 1, because we have 1 room so its no longer blocking

	fmt.Println(<-x) // we are able to pull it off
}
