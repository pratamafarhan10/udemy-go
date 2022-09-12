package main

import "fmt"

func write(channel chan<- string) {
	channel <- "hello there"
}

func main() {
	c := make(chan string)

	go write(c)

	fmt.Println(<-c)
}
