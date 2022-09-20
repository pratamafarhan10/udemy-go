package main

import "fmt"

var x bool

func main() {
	fmt.Println(x) // this will print out false, because false is the zero value of boolean
	x = true
	fmt.Println(x) // this will print out true

	a := 7
	b := 42
	fmt.Println(a == b) // false
	fmt.Println(a <= b) // true
	fmt.Println(a >= b) // false
	fmt.Println(a != b) // true
}
