package hello

import "fmt"

func ExampleSayHello() {
	fmt.Println(SayHello("Sule"))
	// Output:
	// Hello Sule, have a good day!
}

func ExampleSayHello_second() {
	fmt.Println(SayHello("Sule"))
	// Output:
	// Hello Sule, have a good day!
}

func ExamplePerson_Introduction() {
	p1 := Person{
		Name: "Sule",
		Age:  53,
	}

	fmt.Println(p1.Introduction())
	// Output:
	// Hi, my name is Sule and I'm 53 years old
}
