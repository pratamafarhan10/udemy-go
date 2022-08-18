package main

import "fmt"

func main() {
	p1 := struct {
		first, last string
		age         int
	}{
		first: "farhan",
		last:  "pratama",
		age:   22,
	}

	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)
}
