package main

import "fmt"

func main() {
	a := foo()
	b, c := bar()
	fmt.Println(a, b, c)
}

func foo() int {
	return 11
}

func bar() (int, string) {
	return 12, "Hello world"
}
