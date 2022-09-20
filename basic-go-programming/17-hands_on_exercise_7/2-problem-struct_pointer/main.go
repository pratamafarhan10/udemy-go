package main

import "fmt"

type person struct {
	firstName, lastName string
}

func changeMe(p *person) {
	// *p = person{
	// 	firstName: "andre",
	// 	lastName:  "taulany",
	// }

	// *&p.firstName = "andre"
	// *&p.lastName = "andre"
	fmt.Println("value from p address:", *p)
	fmt.Println("value of p:", p)
	fmt.Println("address of p stored:", &p)
	// (*p).firstName = "andre" // dereferencing with pointer in struct
	// (*p).lastName = "taulany"
	p.firstName = "andre"
	p.lastName = "taulany"
}

func main() {
	person1 := person{
		firstName: "sule",
		lastName:  "prikitiw",
	}
	fmt.Println("address of person1.firstName:", &person1.firstName)
	fmt.Println("address of person1.lastName:", &person1.lastName)

	// fmt.Println(person1)

	changeMe(&person1)

	fmt.Println(person1)
}
