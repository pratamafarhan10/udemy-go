package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Goroutine:", runtime.NumGoroutine())
	fmt.Println("CPU:", runtime.NumCPU())

	var counter int64

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Goroutine:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Goroutine:", runtime.NumGoroutine())
	fmt.Println("CPU:", runtime.NumCPU())

	fmt.Println(counter)
}
