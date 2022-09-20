package main

import "fmt"

func main() {
	x := 2
	fmt.Println(x)

	change(&x, 10)
	fmt.Println(x)
}

func change(originalAddress *int, newValue int) {
	*originalAddress = newValue
}
