package main

import "fmt"

func main() {
	foo()
	bar("farhan")
	name := woo("farhan")

	fmt.Println(name)

	message, legalAge := pop("Farhan", 18)

	fmt.Println(message)
	fmt.Println(legalAge)
}

// func (r receiver) identifier(parameter) (return(s)) { ... }
// ketika mendeklarasikan function kalau ada parameter kita sebut parameter
// tetapi ketika kita memanggil function kita memberikan argument bukan parameter

func foo() {
	fmt.Println("Hello from foo")
}

// Everything in GO is pass by value
func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hi, my name is ", s)
}

func pop(firstName string, age int) (string, bool) {
	text := fmt.Sprint("Hi my name is ", firstName, " and I'm ", age, " years old")
	legalAge := false
	if age > 17 {
		legalAge = true
	}

	return text, legalAge
}
