package main

import "fmt"

func main() {
	// x := [5]int{}
	// x := make([]int, 5, 10)
	x := []int{}

	for i := 0; i < 5; i++ {
		// x[i] = i + 1
		x = append(x, i+1)
	}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", x)
}
