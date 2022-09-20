package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)

	// receive
	receive(even, odd, quit)

	fmt.Println("About to quit")
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("even:", v)
		case v := <-odd:
			fmt.Println("odd:", v)
		case v := <-quit:
			fmt.Println("quit:", v)
			return
		}
	}
}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	// close(even)
	// close(odd)

	quit <- 0
}
