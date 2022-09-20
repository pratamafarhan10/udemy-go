package main

import "fmt"

func main() {
	var answer1, answer2, answer3 string

	fmt.Println("What is your name?")
	_, err := fmt.Scan(&answer1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("What is your favorite food?")
	_, err = fmt.Scan(&answer2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("How old are you?")
	_, err = fmt.Scan(&answer3)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(answer1, answer2, answer3)
}
