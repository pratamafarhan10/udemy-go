package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	fmt.Println(incrementor(10, 20))
}

func incrementor(init, to int) int {
	for i := 0; i < to; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			v := init
			v++
			init = v
			defer wg.Done()
			defer mu.Unlock()
		}()
	}
	wg.Wait()
	return init
}
