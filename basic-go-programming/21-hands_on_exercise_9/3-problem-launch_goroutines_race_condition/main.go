package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(incrementor(10, 20))
}

func incrementor(init, to int) int {
	for i := 0; i < to; i++ {
		wg.Add(1)
		go func() {
			v := init
			runtime.Gosched()
			v++
			init = v
			defer wg.Done()
		}()
	}
	wg.Wait()
	return init
}
