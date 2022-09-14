package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	defer foo()

	if err != nil {
		panic(err) // run defer function
	}
}

func foo() {
	fmt.Println("Hello world")
}
