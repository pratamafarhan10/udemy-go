package main

import "fmt"

func main() {
	age := map[string]int{
		"farhan": 22,
		"james":  40,
		"sule":   39,
	}

	fmt.Println(age)

	if v, ok := age["james"]; ok {
		fmt.Println(v)
		delete(age, "james")
	}

	fmt.Println(age)
}
