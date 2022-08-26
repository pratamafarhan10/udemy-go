package main

import "fmt"

func main() {
	numbers := []int{42, 674, 12, 75, 34, 6745, 23}

	sumFoo := foo(numbers...)
	sumBar := bar(numbers)

	fmt.Println("foo", sumFoo)
	fmt.Println("bar", sumBar)
}

func foo(x ...int) int {
	total := 0
	for _, val := range x {
		total += val
	}

	return total
}

func bar(x []int) int {
	total := 0
	for _, val := range x {
		total += val
	}

	return total
}
