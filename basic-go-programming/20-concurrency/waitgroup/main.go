package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
	fmt.Println("Number CPU:", runtime.NumCPU())
	fmt.Println("Number Go routine:", runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	bar()

	fmt.Println("Number CPU:", runtime.NumCPU())
	fmt.Println("Number Go routine:", runtime.NumGoroutine())

	go zee()

	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar", i)
	}
}

func zee() {
	for i := 0; i < 10; i++ {
		fmt.Println("zee", i)
	}
	wg.Done()
}

// * Concurrency is when two or more tasks can start, run, and complete in overlapping time periods. It doesn't necessarily mean they'll ever both be running at the same instant. For example, multitasking on a single-core machine.
// * Parallelism is when tasks literally run at the same time, e.g., on a multicore processor.
// * ketika kita menggunakan syntax "go" maka function tersebut akan berjalan di go routine yang berbeda, akan tetapi ketika function main berhenti maka semua program akan berhenti WALAUPUN si function foo itu baru berjalan atau belum berjalan sama sekali
// * WaitGroup has nothing to do with CPU count. It has to do with making a goroutine wait until other goroutines have finished working. If you are familiar with other programming languages, this is similar to a join on a thread.
// * https://pkg.go.dev/sync#WaitGroup
