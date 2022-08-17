package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"pratama_farhan":  {"Code", "read", "es teh manis"},
	}

	fmt.Println(x)

	x["taulany_andre"] = []string{"Pelawak", "mobil", "ice cream"}

	for i, person := range x {
		fmt.Printf("%v\n", i)
		for _, val := range person {
			fmt.Printf("\t%v\n", val)
		}
	}
}
