package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))
}

func factorial(n int) int {
	x := 1
	for i := n; i > 0; i-- {
		x *= i
	}

	return x
}
