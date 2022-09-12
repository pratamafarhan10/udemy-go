package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for val := range c2 {
		fmt.Println(val)
	}

	fmt.Println("========= About to exit")
}

func populate(c1 chan int) {
	for i := 0; i < 50; i++ {
		c1 <- i
	}

	close(c1)
}

// Throttling with specific number of goroutine
// func fanOutIn(c1, c2 chan int) {
// 	var wg sync.WaitGroup

// 	const goroutine = 10

// 	wg.Add(goroutine)

// 	for i := 0; i < goroutine; i++ {
// 		go func() {
// 			for v := range c1 {
// 				c2 <- timeConsumingWork(v)
// 			}
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	close(c2)
// }

// My own try on Throttling
func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup

	const goroutine = 10

	wg.Add(goroutine)
	for i := 0; i < goroutine; i++ {
		go func() {
			c2 <- timeConsumingWork(<-c1)
			defer wg.Done()
		}()
	}

	wg.Wait()
	close(c2)
}

// No specific number of goroutine
// func fanOutIn(c1, c2 chan int) {
// 	var wg sync.WaitGroup

// 	for val := range c1 {
// 		wg.Add(1)
// 		go func(v2 int) {
// 			c2 <- timeConsumingWork(v2)
// 			defer wg.Done()
// 		}(val)
// 	}

// 	wg.Wait()
// 	close(c2)
// }

func timeConsumingWork(val int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return val + rand.Intn(1000)
}
