package main

import "fmt"

var x string
var y int

func main() {
	fmt.Println("printing the value of x", x, "end")
	fmt.Printf("%T\n", x)

	x = "bond, james bond"
	fmt.Println("printing the value of x", x, "end")
	fmt.Printf("%T\n", x)

	fmt.Println("printing the value of y", y, "end")
	fmt.Printf("%T\n", y)
}
