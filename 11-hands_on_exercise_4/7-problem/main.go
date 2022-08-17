package main

import "fmt"

func main() {
	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooo, james"}}

	for i, people := range x {
		fmt.Println("Perulangan ke", i)
		for _, val := range people {
			fmt.Println(val)
		}
	}
}
