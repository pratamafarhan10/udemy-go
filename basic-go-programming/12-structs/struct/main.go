package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type people struct {
	firstPerson  person
	secondPerson person
}

func main() {
	p1 := person{
		first: "farhan",
		last:  "pratama",
		age:   22,
	}

	p2 := person{
		first: "james",
		last:  "bond",
		age:   42,
	}

	a := people{
		firstPerson:  p1,
		secondPerson: p2,
	}

	fmt.Println(p1)
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(a)
	fmt.Println(a.firstPerson, a.secondPerson)

	// Assign new value
	a.firstPerson = p2
	fmt.Println(a.firstPerson, a.secondPerson)
}
