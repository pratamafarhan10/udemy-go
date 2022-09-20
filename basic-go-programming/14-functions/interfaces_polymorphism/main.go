package main

import "fmt"

type person struct {
	firstName, lastName string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	amIHuman()
}

type cobaInterface interface{}

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

	person1 := person{
		firstName: "Andre",
		lastName:  "Taulany",
	}

	fmt.Println("============== Introduce method")
	fmt.Println(agent1)

	agent1.introduce()
	agent2.introduce()

	fmt.Println("============== Able to kill someone method")
	ltk1 := agent1.amIAbleToKillSomeone()
	ltk2 := agent2.amIAbleToKillSomeone()

	fmt.Println(ltk1)
	fmt.Println(ltk2)

	fmt.Println("============== Am i human method")
	agent1.amIHuman()
	agent2.amIHuman()
	person1.amIHuman()
	// person1.introduce() // this will throw an error because the method introduce only attached to type secretAgent

	fmt.Println("============== Interface and polymorphism")
	bar(agent1)
	bar(agent2)
	bar(person1)

	fmt.Println("============== Empty interface")
	emptyInterface("farhan")
	emptyInterface(2090)
	emptyInterface(agent2)
	emptyInterface(person1)

	fmt.Println("============== Interface variadic parameter")
	x := []int{1, 2, 3, 4, 5}
	printAll(1, "farhan", 3, false, true, x)

	fmt.Println("============== Empty interface as a parameter type")
	noType("farhan")
	noType(2090)
	noType(agent2)
	noType(person1)
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

func (p person) amIHuman() {
	fmt.Println("yes,", p.firstName, "you are human")
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("this is a person", h.(person).firstName)
	case secretAgent:
		fmt.Println("this is a secret agent", h.(secretAgent).firstName)
	}

	fmt.Println("i was passed into bar", h)
}

func emptyInterface(a cobaInterface) {
	switch a.(type) {
	case person:
		fmt.Println("this is a person", a.(person).firstName)
	case secretAgent:
		fmt.Println("this is a secret agent", a.(secretAgent).firstName)
	case int:
		fmt.Println("this is a integer", a)
	case string:
		fmt.Println("this is a string", a)
	}
}

func printAll(data ...interface{}) {
	fmt.Println("the length of data:", len(data))
	fmt.Printf("%T\t", data...)
	fmt.Println(data...)
}

func noType(a interface{}) {
	switch a.(type) {
	case person:
		fmt.Println("this is a person", a.(person).firstName)
	case secretAgent:
		fmt.Println("this is a secret agent", a.(secretAgent).firstName)
	case int:
		fmt.Println("this is a integer", a)
	case string:
		fmt.Println("this is a string", a)
	}
}
