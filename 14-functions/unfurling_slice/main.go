package main

import "fmt"

func main() {
	fmt.Println("=============== sum 1")
	x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Sum of x:", x)

	fmt.Println("=============== bar")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	bar(numbers)

	fmt.Println("=============== sum 2")
	y := sum(numbers...)
	fmt.Println("Sum of y:", y)

	fmt.Println("=============== average")
	z := average(9, numbers...)
	// z := average(9)
	fmt.Println("Average of z:", z)

	// fmt.Println("=============== foo")
	// z := foo() // this wouldn't work because we didn't pass any arguments
	// fmt.Println(z)
}

func sum(x ...int) int { // Variadic means we can pass 0 or more arguments AND it has to be the final parameter

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for _, val := range x {
		sum += val
	}

	return sum
}

func average(count int, x ...int) int { // Variadic means we can pass 0 or more arguments AND it has to be the final parameter

	fmt.Println(count, x)
	fmt.Printf("%T %T\n", count, x)

	sum := 0

	for _, val := range x {
		sum += val
	}

	return sum / count
}

// func foo(s string) string {
// 	return s
// }

func bar(x []int) {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for _, val := range x {
		sum += val
	}

	fmt.Println("Sum of x:", sum)
}
