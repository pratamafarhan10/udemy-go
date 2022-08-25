package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("this is a func expression")
	}

	a := func(x int) {
		fmt.Println("the value of x:", x)
	}

	sum := func(a, b int) int {
		return a + b
	}

	f()
	a(2)
	fmt.Println("sum of 2 and 3:", sum(2, 3))
}
