package main

import "fmt"

func main() {
	x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Sum of x:", x)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	bar(numbers)
}

func sum(x ...int) int {
	fmt.Println("=============== sum")
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for _, val := range x {
		sum += val
	}

	return sum
}

func bar(x []int) {
	fmt.Println("=============== bar")
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for _, val := range x {
		sum += val
	}

	fmt.Println("Sum of x:", sum)
}
