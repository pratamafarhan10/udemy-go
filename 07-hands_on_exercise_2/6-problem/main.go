package main

import "fmt"

const (
	a = 2022 - (iota + 1)
	b
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)
}
