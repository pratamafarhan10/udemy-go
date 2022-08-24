package main

import "fmt"

type person struct {
	firstName, lastName string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func main() {
	agent1 := secretAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
		},
		licenseToKill: true,
	}

	fmt.Println(agent1)

	agent1.introduce()
}

func (s secretAgent) introduce() {
	fmt.Println("My name is", s.firstName, s.lastName)
}

// func (r receiver) identifier(parameters) (return/returns){ ... }
