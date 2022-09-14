package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	defer foo()

	if err != nil {
		log.Panicln(err) // run defer function
	}
}

func foo() {
	fmt.Println("Hello world")
}
