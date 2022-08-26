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

	people := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, val := range people {
		fmt.Println(val.firstName, val.lastName, "favorite flavors of ice cream is:")
		for _, flavor := range val.favoriteFlavor {
			fmt.Printf("\t%v\n", flavor)
		}
		fmt.Println("---------------------")
	}

}
