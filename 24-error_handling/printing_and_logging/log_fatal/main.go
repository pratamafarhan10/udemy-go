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
		log.Fatalln(err) // doesn't run defer function, straight to exit
	}
}

func foo() {
	fmt.Println("Hello")
}
