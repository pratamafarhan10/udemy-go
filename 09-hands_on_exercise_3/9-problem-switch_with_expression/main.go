package main

import "fmt"

func main() {
	favSport := "football"

	switch favSport {
	case "american football":
		fmt.Println("this is not a sport")
	case "football":
		fmt.Println("football is my favorite sport")
	default:
		fmt.Println("i don't have any favorite sport")
	}
}
