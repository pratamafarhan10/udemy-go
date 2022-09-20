package main

import (
	"fmt"
	"math/rand"
)

func main() {
	evenChannel := make(chan string)
	oddChannel := make(chan string)

	randNumbers := generateRandomNumber(10)

	go pullOut(evenChannel)
	go pullOut(oddChannel)

	fanOut(evenChannel, oddChannel, randNumbers)
	fmt.Println("About to exit")
}

// generateRandomNumber generates n random number from 0-100
func generateRandomNumber(length int) chan int {
	rndm := make(chan int)
	go func(lgth int) {
		for i := 0; i < lgth; i++ {
			rndm <- rand.Intn(100)
		}

		close(rndm)
	}(length)

	return rndm
}

// fanOut sort the value into even or odd channel
func fanOut(evenChannel, oddChannel chan<- string, randNumber chan int) {
	for val := range randNumber {
		if val%2 == 0 {
			evenChannel <- fmt.Sprint("even:", val)
		} else {
			oddChannel <- fmt.Sprint("odd:", val)
		}
	}
}

// pullOut pulls values from channel
func pullOut(c <-chan string) {
	for val := range c {
		fmt.Println(val)
	}
}
