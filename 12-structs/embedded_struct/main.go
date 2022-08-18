package main

import "fmt"

type person struct {
	// first, last string // we can do this or the bottom one
	first string
	last  string
	age   int
}

type secretAgent struct {
	person        // fieldname will be person
	first         string
	licenseToKill bool
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

	a := secretAgent{
		// person: person{
		// 	first: "farhan",
		// 	last:  "pratama",
		// 	age:   22,
		// },
		person:        p1, // we can do this or above
		first:         "collision",
		licenseToKill: false,
	}

	fmt.Println("========================= Default struct")
	fmt.Println(p1)
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println("========================= Embedded struct")
	fmt.Println(a)
	fmt.Println(a.first, a.person.first, a.last, a.age, a.licenseToKill)
}
