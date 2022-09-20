package main

import "fmt"

func main() {
	// A slice allows you to group together values of the same type
	// Composite literals
	x := []int{4, 12, 623, 21, 84}
	fmt.Println(x)
	fmt.Println(x[:])   // the whole array
	fmt.Println(x[1:])  // give array from index 1 to end
	fmt.Println(x[1:3]) // give array from index 1 to 2

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
