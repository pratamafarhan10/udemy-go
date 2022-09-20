package main

import "fmt"

func main() {
	fmt.Println("Hello, World! We are currently learning about the Go programming language.")

	foo()

	fmt.Println("We are done with the foo function.")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("We are in the foo function.")
}

func bar() {
	fmt.Println("and we are exited")
}
