package main

import "fmt"

func main() {
	fmt.Println("=========== Foo")
	foo()

	fmt.Println("=========== Func expression with return")
	a := func() int {
		return 451
	}

	fmt.Printf("%T\n", a)

	fmt.Println("=========== Func returning a func")
	b := bar()

	fmt.Printf("%T\n", b)

	fmt.Println("=========== Execute the func")
	c := b()

	fmt.Printf("type: %T, return: %v\n", b, c)

	fmt.Println("=========== Execute the func with two parentheses")
	fmt.Println(bar()())
}

func foo() {
	fmt.Println("Hello world")
}

func bar() func() int {
	return func() int {
		return 451
	}
}
