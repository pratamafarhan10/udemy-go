package hello

import "fmt"

type Person struct {
	Name string
	Age  int
}

func SayHello(name string) string {
	return fmt.Sprintf("Hello %v, have a good day!", name)
}

func (p Person) Introduction() string {
	return fmt.Sprintf("Hi, my name is %v and I'm %v years old", p.Name, p.Age)
}
