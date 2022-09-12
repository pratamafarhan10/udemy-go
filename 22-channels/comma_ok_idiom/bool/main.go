package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send(even, odd, quit)

	// receive
	receive(even, odd, quit)
	fmt.Println("About to quit")
}

func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	// quit <- true
	close(quit)
}

func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("Even:", v)
		case v := <-odd:
			fmt.Println("Odd:", v)
		case v, open := <-quit: // Check whether the channel is still open or not
			if !open { // the channel is close
				fmt.Printf("Quit. v = %v, open: %v\n", v, open)
				return
			} else {
				fmt.Println("Quit. v:", v)
				return
			}
		}
	}
}
