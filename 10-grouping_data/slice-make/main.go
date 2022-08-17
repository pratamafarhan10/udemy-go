package main

import "fmt"

func main() {
	x := make([]int, 10, 12)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 1
	x[9] = 10

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// x[11] = 11 // return error because run out of length
	x = append(x, 11)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 12, 13)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
