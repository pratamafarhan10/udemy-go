package main

import (
	"fmt"
	"sync"
)

func main() {
	names := []string{"andre", "taulany", "adul", "future", "drake", "sule", "komeng"}

	fmt.Println(renameWithoutConcurrency(names))
	fmt.Println(renameWithConcurrency(names))
}

func renameWithoutConcurrency(names []string) []string {
	for i := range names {
		names[i] = "james bond"
	}

	return names
}

func renameWithConcurrency(names []string) []string {
	var wg sync.WaitGroup

	for i := range names {
		wg.Add(1)
		go func(n int) {
			names[n] = "james bond"
			defer wg.Done()
		}(i)
	}

	wg.Wait()

	return names
}
