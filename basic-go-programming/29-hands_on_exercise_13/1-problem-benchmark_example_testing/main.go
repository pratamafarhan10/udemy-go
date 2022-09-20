package main

import (
	"fmt"

	"github.com/udemy-go/29-hands_on_exercise_13/1-problem-benchmark_example_testing/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
