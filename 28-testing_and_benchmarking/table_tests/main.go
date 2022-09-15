package main

import "fmt"

func main() {
	fmt.Println(mySum(1, 2, 3))
	fmt.Println(mySum(21, 21))
	fmt.Println(mySum(42, 23, 41))
}

func mySum(n ...int) int {
	x := 0

	for _, val := range n {
		x += val
	}

	return x
}
