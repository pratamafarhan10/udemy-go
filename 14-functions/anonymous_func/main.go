package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("anonymous function")
	}()

	func(x int) {
		fmt.Println("the value of x:", x)
	}(42)
}

func foo() {
	fmt.Println("foo runs")
}
