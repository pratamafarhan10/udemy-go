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

// *I would completely consider it best practice to use defer for the unlock. Go is very good at optimizing defer statements, so it won't cost you any speed. The only time I would skip defer for unlock, is if I am locking, running a statement that has zero chances of panicing, and then unlocking.
