package main

import "fmt"

type person struct {
	firstName, lastName string
	age                 int
}

func main() {
	person1 := person{
		firstName: "Andre",
		lastName:  "Taulany",
		age:       51,
	}

	person1.speak()
}

func (p person) speak() {
	fmt.Println("Hi, my name is", p.firstName, p.lastName, "and I'm", p.age, "years old")
}
