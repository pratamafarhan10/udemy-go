package main

import "fmt"

func main() {
	c := make(chan int)
	cs := make(chan<- int)
	cr := make(<-chan int)

	fmt.Println("Print type")
	fmt.Printf("non directional: %T\n", c)
	fmt.Printf("receive: %T\n", cr)
	fmt.Printf("send: %T\n", cs)

	fmt.Println("\nAssign works: general to specific")
	cs = c
	cr = c

	// fmt.Println("Assign doesn't works: specific to general")
	// c = cs
	// c = cr

	fmt.Println("\nConversion works: general to specific")
	fmt.Printf("receive: %T\n", (<-chan int)(c))
	fmt.Printf("send: %T\n", (chan<- int)(c))

	// fmt.Println("Conversion doesn't works: specific to general")
	// fmt.Printf("receive: %T\n", (chan int)(cr))
	// fmt.Printf("send: %T\n", (chan int)(cs))
}
