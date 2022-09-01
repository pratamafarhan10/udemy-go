package library

import "fmt"

type Student struct {
	Name  string
	Grade int
}

var Student1 Student

func init() {
	Student1 = Student{
		Name:  "Sule",
		Grade: 99,
	}

	fmt.Println("-> package library imported")
}

func SayHello(name string) {
	fmt.Println("Hello")
	introduce(name)
}

func introduce(name string) {
	fmt.Println("nama saya:", name)
}
