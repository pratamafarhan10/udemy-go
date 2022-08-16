package main

import "fmt"

func main() {
	x := []int{4, 12, 623, 21, 84}
	fmt.Println(x)
	x = append(x, 41, 2123, 6423, 1)
	fmt.Println(x)

	y := []int{5623, 123, 643}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:5], x[7:]...)
	fmt.Println(x)
}
