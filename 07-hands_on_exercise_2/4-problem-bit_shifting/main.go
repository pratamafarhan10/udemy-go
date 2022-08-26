package main

import "fmt"

func main() {
	// assign an int to a variable
	a := 20

	// print in decimal, binary, and hexadecimal
	fmt.Printf("%d\t\t%b\t\t%#x\n", a, a, a)

	// shift the bits over 1 position to the left and assign it to a variable
	b := 20 << 1

	// print in decimal, binary, and hexadecimal
	fmt.Printf("%d\t\t%b\t\t%#x\n", b, b, b)
}
