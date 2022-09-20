package main

import "fmt"

type person struct {
	firstName, lastName string
	age                 int
}

func (p *person) speak() {
	fmt.Println("Hi, my name is", p.firstName, p.lastName, "and I'm", p.age, "years old.")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		firstName: "Andre",
		lastName:  "Taulany",
		age:       45,
	}

	fmt.Println("============= Call the function speak directly")
	p1.speak() // this will run because the compiler will translate it to this (&p1).speak()

	fmt.Println("============= Call the function speak with ampersand")
	(&p1).speak()

	// fmt.Println("============= Store the value p1 to a variable with a type human")
	// var h1 human
	// h1 = p1 // this will throw an error because a type human with a method speak only receive a pointer type

	fmt.Println("============= Store the value p1 with a pointer type to a variable with a type human")
	var h1 human
	h1 = &p1
	fmt.Println(h1)
	h1.speak()

	// fmt.Println("============= Call the function saySomething by pass a non pointer argument")
	// saySomething(p1) // this will throw an error because type human with function speak only receive a pointer type

	fmt.Println("============= Call the function saySomething by pass a pointer argument")
	saySomething(&p1)
}
