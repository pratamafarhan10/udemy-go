package main

import (
	"fmt"

	"github.com/udemy-go/basic-go-programming/28-testing_and_benchmarking/example_tests/hello"
)

func main() {
	fmt.Println("======================= Testing a function")
	fmt.Println(hello.SayHello("Farhan"))
	fmt.Println(hello.SayHello("Todd"))
	fmt.Println(hello.SayHello("James Bond"))

	fmt.Println("======================= Testing a method of a type")
	p1 := hello.Person{
		Name: "Andre",
		Age:  53,
	}

	fmt.Println(p1.Introduction())

}
