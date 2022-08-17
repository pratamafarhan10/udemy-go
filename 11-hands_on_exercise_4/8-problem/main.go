package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"pratama_farhan":  {"Code", "read", "es teh manis"},
	}

	fmt.Println(x)

	for i, person := range x {
		fmt.Printf("This is the %v loop\n", i)
		for _, val := range person {
			fmt.Printf("\t%v\n", val)
		}
	}
}
