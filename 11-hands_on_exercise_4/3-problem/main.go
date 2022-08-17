package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(x)

	// 42 - 46
	fmt.Println(x[:5])
	// 47 - 51
	fmt.Println(x[5:])
	// 44 - 48
	fmt.Println(x[2:7])
	// 43 - 47
	fmt.Println(x[1:6])
}
