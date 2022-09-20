package main

import "fmt"

func main() {
	a := 42

	fmt.Println("============= Show the value and the address")
	fmt.Println("the value of a:", a)
	fmt.Println("the address of a:", &a) // to show the address where the value a is stored. & gives you the address of the value

	fmt.Println("============= Show the type of the value and the address")
	fmt.Printf("the type of a: %T\n", a)
	fmt.Printf("the type of the address of a: %T\n", &a) // show the type of the address of a (*int), and it's totally different with int

	fmt.Println("============= Show the address and the type of the value after store it into another variable")
	var b *int = &a // the address of a with the type of *int is stored in variable b
	fmt.Println("the value of b:", b)
	fmt.Println("the address of a:", &b)
	fmt.Printf("the type of b: %T\n", b)

	fmt.Println("============= Dereferencing an address")
	fmt.Println("the value of b:", *b) // * gives the value stored in the address

	fmt.Println("============= Reassign the value based on address")
	*b = 43
	fmt.Println("the value of a:", a)
}
