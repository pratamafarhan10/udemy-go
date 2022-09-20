package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Goroutine:", runtime.NumGoroutine())
	fmt.Println("CPU:", runtime.NumCPU())

	counter := 0
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutine:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Goroutine:", runtime.NumGoroutine())
	fmt.Println("CPU:", runtime.NumCPU())

	fmt.Println(counter)
}
