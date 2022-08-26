package main

import "fmt"

func main() {
	numbers := []int{42, 63, 1, 2, 53, 7, 32, 64, 234, 123, 4623, 75, 2, 21}

	sortedNumbers := evenNumbersCheck(sortEvenNumbers, numbers...)

	fmt.Println(sortedNumbers(2))
	fmt.Println(sortedNumbers(10))
}

func sortEvenNumbers(numbers ...int) []int {
	even := []int{}
	for _, val := range numbers {
		if val%2 == 0 {
			even = append(even, val)
		}
	}

	return even
}

func evenNumbersCheck(f func(numbers ...int) []int, numbers ...int) func(number int) bool {
	sortedNumbers := f(numbers...)

	return func(number int) bool {
		for _, val := range sortedNumbers {
			if number == val {
				return true
			}
		}

		return false
	}
}
