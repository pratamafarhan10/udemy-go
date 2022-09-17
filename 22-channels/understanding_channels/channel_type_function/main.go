package main

import "fmt"

func main() {
	c := make(chan func())
	d := make(chan bool)

	go func() {
		c <- one
		d <- true
	}()

	select {
	case v := <-c:
		v()
	case _, ok := <-d:
		if !ok {
			return
		}
	}
}

func one() {
	fmt.Println("1")
}
