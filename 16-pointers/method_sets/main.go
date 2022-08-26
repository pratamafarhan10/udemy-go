package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func foo(s shape) {
	fmt.Println(s.area())
}

func main() {
	// Method sets menentukan method apa saja yang bisa di-attach ke suatu type data
	// receiver (t T) valuesnya bisa T and *t
	// receiver(t *T) valuesnya harus *T
	square1 := square{
		side: 4,
	}

	circle1 := circle{
		radius: 12,
	}

	// foo(square1)
	// foo(circle1)

	fmt.Println("======= If the receiver type is not a pointer we can send the value or the address")
	foo(square1)
	foo(&square1)

	fmt.Println("======= If the receiver type is a pointer we can only send the address")
	foo(&circle1)
	// foo(circle1) // this will throw an error because the function only receive a pointer type
}
