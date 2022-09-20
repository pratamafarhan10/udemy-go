package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(incrementor(10, 20))
}

func incrementor(init int64, to int) int64 {
	for i := 0; i < to; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&init, 1)
			runtime.Gosched()
			defer wg.Done()
		}()
	}
	wg.Wait()
	return init
}
