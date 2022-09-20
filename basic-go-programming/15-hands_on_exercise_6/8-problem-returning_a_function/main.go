package main

import "fmt"

func main() {
	even := findTheEvenNumber(42, 63, 1, 2, 53, 7, 32, 64, 234, 123, 4623, 75, 2, 21)

	fmt.Println(even(2))
	fmt.Println(even(3))
	fmt.Println(even(64))
	fmt.Println(even(234))
	fmt.Println(even(142))
}

func findTheEvenNumber(numbers ...int) func(number int) bool {
	evenNumbers := []int{}
	for _, val := range numbers {
		if val%2 == 0 {
			evenNumbers = append(evenNumbers, val)
		}
	}

	return func(number int) bool {
		for _, val := range evenNumbers {
			if val == number {
				return true
			}
		}
		return false
	}
}
