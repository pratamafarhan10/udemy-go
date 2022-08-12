package main

import "fmt"

type hotdog int

var x hotdog
var y int

func main() {
	// a. problem 4
	// print out the value of x
	fmt.Println(x)

	// print out the type of the variable x
	fmt.Printf("%T\n", x)

	// assign 42 to the variabel x using the "=" operator
	x = 42

	// print out the value of the variable x
	fmt.Println(x)

	// b. problem 5
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
