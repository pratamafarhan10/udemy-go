package main

import (
	"fmt"

	"github.com/udemy-go/27-hands_on_exercise_12/dog"
)

func main() {
	fmt.Println("1 human years is equal to", dog.Years(1), "dog years")
	fmt.Println("2 human years is equal to", dog.Years(2), "dog years")
	fmt.Println("3 human years is equal to", dog.Years(3), "dog years")
	fmt.Println("4 human years is equal to", dog.Years(4), "dog years")
}
