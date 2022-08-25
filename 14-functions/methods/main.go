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

	agent2 := secretAgent{
		person: person{
			firstName: "Miss",
			lastName:  "Moneypenny",
		},
		licenseToKill: false,
	}

	fmt.Println(agent1)

	agent1.introduce()
	agent2.introduce()

	ltk1 := agent1.amIAbleToKillSomeone()
	ltk2 := agent2.amIAbleToKillSomeone()

	fmt.Println(ltk1)
	fmt.Println(ltk2)
}

// func (r receiver) identifier(parameters) (return/returns){ ... }

func (s secretAgent) introduce() {
	fmt.Println("My name is", s.firstName, s.lastName)
}

func (s secretAgent) amIAbleToKillSomeone() string {
	if s.licenseToKill {
		return "you are able to kill someone"
	}

	return "you are not able to kill someone"
}
