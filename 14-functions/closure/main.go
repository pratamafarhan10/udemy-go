package main

import "fmt"

var x int

func main() {
	fmt.Println("============== Package level scope variable")
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)

	fmt.Println("============== New block of code inside main function")

	{
		y := 42
		fmt.Println(y)
	}
	// fmt.Println(y) // This will throw an error because y is inside a different block of code

	fmt.Println("============== Closure")
	a := incrementor()
	b := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func foo() {
	x++
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
