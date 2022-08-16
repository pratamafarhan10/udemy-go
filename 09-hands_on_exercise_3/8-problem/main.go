package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this is false")
	case !true:
		fmt.Println("this is false")
	case 2 == 2:
		fmt.Println("this is true")
	default:
		fmt.Println("default")
	}
}
