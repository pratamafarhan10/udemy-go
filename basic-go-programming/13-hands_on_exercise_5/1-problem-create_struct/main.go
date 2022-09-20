package main

import "fmt"

type person struct {
	firstName, lastName string
	favoriteFlavor      []string
}

func main() {
	p1 := person{
		firstName:      "farhan",
		lastName:       "pratama",
		favoriteFlavor: []string{"strawberry", "chocolatte", "blueberry"},
	}

	p2 := person{
		firstName:      "james",
		lastName:       "bond",
		favoriteFlavor: []string{"hazelnut", "vanilla", "oreo"},
	}

	for i, val := range p1.favoriteFlavor {
		fmt.Println(i, val)
	}

	for i, val := range p2.favoriteFlavor {
		fmt.Println(i, val)
	}
}
