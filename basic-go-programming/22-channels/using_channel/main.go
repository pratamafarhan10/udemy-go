package main

import "fmt"

func main() {

	c := make(chan string)

	go write(c)

	fmt.Println(read(c))
}

func write(channel chan<- string) {
	channel <- "Hello there"
}

func read(channel <-chan string) string {
	return <-channel
}
