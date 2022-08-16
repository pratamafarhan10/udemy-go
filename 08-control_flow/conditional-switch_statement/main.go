package main

import "fmt"

func main() {
	// switch {
	// case false:
	// 	fmt.Println("this should not print")
	// case 2 == 4:
	// 	fmt.Println("this should not print")
	// case 3 == 3:
	// 	fmt.Println("print")
	// // fallthrough // to run the rest of the statement (even if it's false)
	// // case 4 == 4:
	// // 	fmt.Println("also true, does it print?")
	// // 	fallthrough
	// // case 6 == 15:
	// // 	fmt.Println("not true 1")
	// // 	fallthrough
	// // case 7 == 15:
	// // 	fmt.Println("not true 1")
	// // 	fallthrough
	// // case 15 == 15:
	// // 	fmt.Println("true 15")
	// default:
	// 	fmt.Println("default")
	// }

	x := "james"

	switch x {
	case "james", "bond":
		fmt.Println("james bond")
	case "farhan":
		fmt.Println("farhan")
	case "todd":
		fmt.Println("todd")
	default:
		fmt.Println("default")
	}
}
