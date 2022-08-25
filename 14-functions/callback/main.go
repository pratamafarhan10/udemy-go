package main

import "fmt"

func main() {
	fmt.Println("============== Sum of even number")
	x := sumOfEvenNumbers(sum, []int{1, 2, 3, 6, 1, 4, 7, 2, 7, 3, 57, 84}...)
	fmt.Println(x)

	fmt.Println("============== Sum of odd number")
	y := sumOfOddNumbers(sum, []int{1, 2, 3, 6, 1, 4, 7, 2, 7, 3, 57, 84}...)
	fmt.Println(y)
}

func sum(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func sumOfEvenNumbers(f func(numbers ...int) int, numbers ...int) int {
	evenNumbers := []int{}

	for _, val := range numbers {
		if val%2 == 0 {
			evenNumbers = append(evenNumbers, val)
		}
	}

	return f(evenNumbers...)
}

func sumOfOddNumbers(f func(numbers ...int) int, numbers ...int) int {
	oddNumbers := []int{}

	for _, val := range numbers {
		if val%2 != 0 {
			oddNumbers = append(oddNumbers, val)
		}
	}

	return f(oddNumbers...)
}
