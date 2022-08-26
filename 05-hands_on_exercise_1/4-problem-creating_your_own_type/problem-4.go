package main

import "fmt"

type hotdog int

var x hotdog

func main() {
	// print out the value of x
	fmt.Println(x)

	// print out the type of the variable x
	fmt.Printf("%T\n", x)

	// assign 42 to the variabel x using the "=" operator
	x = 42

	// print out the value of the variable x
	fmt.Println(x)
}
