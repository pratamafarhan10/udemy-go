package main

import "fmt"

// write a type constant and untype constant
const (
	b     = 90.9012
	a int = 42
)

func main() {
	fmt.Println(a, b)
	fmt.Printf("%T\t%T", a, b)
}
