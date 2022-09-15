package hello

import "fmt"

func ExampleSayHello() {
	fmt.Println(SayHello("Farhan"))
	// Output:
	// Hello Farhan, have a good day!
}

func ExampleSayHello_second() {
	fmt.Println(SayHello("Andre"))
	// Output:
	// Hello Andre, have a good day!
}

func ExamplePerson_Introduction() {
	p1 := Person{
		Name: "Andre",
		Age:  53,
	}

	fmt.Println(p1.Introduction())
	// Output:
	// Hi, my name is Andre and I'm 53 years old
}
