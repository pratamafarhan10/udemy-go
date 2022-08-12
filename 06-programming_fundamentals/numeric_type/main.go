package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64
var z int8

func main() {
	x = 42
	y = 32.90123
	// x = 9.0123 this will cause error because x is an integer

	z = -128 // this works
	// z = -129 // this doesn't works
	// z = 127 // this works
	// z = 128 // this doesn't works

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

	// runtime package
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.Version())
}
