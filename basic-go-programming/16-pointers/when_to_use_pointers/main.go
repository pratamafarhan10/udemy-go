package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x before:", x)
	fmt.Println("x address before:", &x)
	foo(&x)
	fmt.Println("x after:", x)
	fmt.Println("x address after:", &x)
}

func foo(x *int) {
	fmt.Println("x foo before:", *x)
	fmt.Println("x foo address before:", x)
	*x = 42
	fmt.Println("x foo after:", *x)
	fmt.Println("x foo address after:", x)
}
