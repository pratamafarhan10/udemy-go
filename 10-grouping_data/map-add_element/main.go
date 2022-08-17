package main

import "fmt"

func main() {
	age := map[string]int{
		"farhan": 22,
		"james":  42,
		"bond":   20,
	}

	fmt.Println(age)

	age["sule"] = 39

	for i, v := range age {
		fmt.Println(i, v)
	}
}
