package main

import "fmt"

func main() {
	fmt.Println("============ Recursive factorial")
	fmt.Println(factorial(4))

	fmt.Println("============ Loop factorial")
	fmt.Println(factorialLoop(4))
}

func factorial(number int) int {
	if number == 1 {
		return 1
	}

	return number * factorial(number-1)
}

func factorialLoop(number int) int {
	total := 1
	for i := number; i >= 1; i-- {
		total *= i
	}

	return total
}
