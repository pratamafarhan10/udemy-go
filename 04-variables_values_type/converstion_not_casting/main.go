package main

import "fmt"

var a int

type hotdog int // we can create our own type

var b hotdog // assign variable with the type hotdog

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 90
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// a = b this wouldn't work because a and b has a different type
	// but we can convert it with this
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T", a)
}
