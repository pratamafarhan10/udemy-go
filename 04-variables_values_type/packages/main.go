package main

import "fmt"

func main() {
	const name, age = "farhan", 20

	// This is how to catch the return parameter from Println function
	n, err := fmt.Println("Hi my name is", name, "and i'm", age, "years old.")
	fmt.Println(n)
	fmt.Println(err)

	// This is how to ignore the error parameter or the number of byte parameter
	// n, _ := fmt.Println("Hi my name is", name, "and i'm", age, "years old.")
	// fmt.Println(n)

	// _, err := fmt.Println("Hi my name is", name, "and i'm", age, "years old.")
	// fmt.Println(err)
}
