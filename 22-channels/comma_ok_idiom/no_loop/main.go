package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
		close(c)
	}()

	v, open := <-c

	fmt.Println(v, open)

	v, open = <-c

	fmt.Println(v, open)
}
