package main

import "fmt"

func main() {

	result := factorial(5)

	fmt.Println(result)

}

func factorial(i int) int {
	if i == 1 {
		return 1
	}

	return i * factorial(i-1)
}
