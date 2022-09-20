package main

import (
	"fmt"
)

func main() {
	a := "Hello world"
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Printf("%T\n", a)

	b := []byte(a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	for i := 0; i < len(a); i++ {
		fmt.Printf("%#U\n", a[i])
	}

	for i, v := range a {
		fmt.Printf("at index position of %d we have hex %#x\n", i, v)
	}
}
