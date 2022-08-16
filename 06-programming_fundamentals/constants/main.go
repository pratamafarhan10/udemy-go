package main

import "fmt"

const (
	a int     = 42
	b float64 = 90.82
	c string  = "james bond"
)

const (
	d = 12
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n%T\n%T\n", a, b, c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
