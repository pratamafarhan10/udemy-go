package main

import "fmt"

func main() {
	increment := incrementor(20)

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}

func incrementor(number int) func() int {
	x := number

	return func() int {
		x++
		return x
	}
}
