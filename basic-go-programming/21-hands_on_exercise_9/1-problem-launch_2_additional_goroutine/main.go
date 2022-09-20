package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	fmt.Println("Num of goroutine", runtime.NumGoroutine())
	fmt.Println("============= Start")
	counter := 0
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			counter++
			fmt.Println("Counter", counter)
			defer wg.Done()
			defer mu.Unlock()
		}()
		fmt.Println("\tNum of goroutine", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("============= Done")
	fmt.Println("Num of goroutine", runtime.NumGoroutine())
}
