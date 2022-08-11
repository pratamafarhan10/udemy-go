package main

import "fmt"

var y = 90 // the scope is for the whole package

var z int // define a zero value variable

func main() {
	// declare a VARIABLE and assign a VALUE
	x := 42 /// the scope only in main function
	fmt.Println("x = ", x)

	fmt.Println("first y = ", y)

	z = 120

	fmt.Println("first z = ", z)

	foo()
	bar()
}

func foo() {
	fmt.Println("second y:", y)
}

func bar() {
	z = 999
	fmt.Println("second z:", z)
}

// it's better to use short declaration operator when you define a variable inside a function
// use a var keyword to define a variable outside of a function
