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
	// wg.Add(100)

	var mu sync.Mutex

	// for i := 0; i < 100; i++ {
	// 	go func() {
	// 		mu.Lock()
	// 		v := counter
	// 		runtime.Gosched()
	// 		v++
	// 		counter = v
	// 		mu.Unlock()
	// 		wg.Done()
	// 	}()
	// 	fmt.Println("Goroutine:", runtime.NumGoroutine())
	// }

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer wg.Done()
			counter++
			defer mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println("Goroutine:", runtime.NumGoroutine())
	fmt.Println("CPU:", runtime.NumCPU())

	fmt.Println(counter)
}
