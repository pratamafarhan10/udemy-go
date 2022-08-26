package main

import "fmt"

func main() {
	fmt.Println("Hello")
	defer foo()
	fmt.Println("World")
}

func foo() {
	fmt.Println("This is deferred func")
}
