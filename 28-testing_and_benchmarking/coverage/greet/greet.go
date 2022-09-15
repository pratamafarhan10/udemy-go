package greet

import "fmt"

func Greet(name string) string {
	return fmt.Sprint("Welcome my dear ", name)
}

func Introduction(name string) string {
	return fmt.Sprint("Hi, my name is ", name)
}
